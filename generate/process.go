package generate

import (
	"fmt"
	"strings"

	"github.com/iancoleman/strcase"
	"github.com/inigolabs/fezzik/common"
	"github.com/inigolabs/fezzik/config"
	"github.com/inigolabs/fezzik/fezzik_ast"
	"github.com/vektah/gqlparser/v2/ast"
	"golang.org/x/tools/go/packages"
)

var mode = packages.NeedName | packages.NeedTypes | packages.NeedModule | packages.NeedTypesInfo

func Process(cfg *config.Config, schema *ast.Schema, operations *ast.QueryDocument) *fezzik_ast.Document {
	result := &fezzik_ast.Document{
		InputTypes:   make(map[string]fezzik_ast.InputType),
		BoundGoTypes: make(map[string]string),
		EnumTypes:    make(map[string]fezzik_ast.EnumType),
		Imports:      make(map[string]string),
	}

	var pkgs []*packages.Package
	var err error
	if len(cfg.Autobind) > 0 {
		pkgs, err = packages.Load(&packages.Config{Mode: mode}, cfg.Autobind...)
		check(err)
	}

	for _, t := range schema.Types {
		switch t.Kind {
		case ast.Enum:
			var found bool
			for _, pkg := range pkgs {
				if foundType := pkg.Types.Scope().Lookup(t.Name); foundType != nil {
					result.BoundGoTypes[t.Name] = foundType.Pkg().Name() + "." + foundType.Name()
					result.Imports[foundType.Pkg().Path()] = ""
					found = true
				}
			}

			if !found {
				result.EnumTypes[t.Name] = fezzik_ast.EnumType{
					Name:   t.Name,
					Values: getEnumValues(t),
				}
			}

		case ast.InputObject:
			var found bool
			for _, pkg := range pkgs {
				if foundType := pkg.Types.Scope().Lookup(t.Name); foundType != nil {
					result.BoundGoTypes[t.Name] = foundType.Pkg().Name() + "." + foundType.Name()
					result.Imports[foundType.Pkg().Path()] = ""
					found = true
				}
			}

			if !found {
				result.InputTypes[t.Name] = fezzik_ast.InputType{
					Name:   t.Name,
					Fields: getInputFields(cfg, result, t),
				}
			}
		case ast.Object:
			for _, pkg := range pkgs {
				if foundType := pkg.Types.Scope().Lookup(t.Name); foundType != nil {
					result.BoundGoTypes[t.Name] = foundType.Pkg().Name() + "." + foundType.Name()
					result.Imports[foundType.Pkg().Path()] = ""
				}
			}
		}
	}

	// walked through defined scalars and collects imports with aliases and scalars full typeNames
	names := map[string]string{}
	for scalarName, path := range cfg.ScalarTypeMap {
		if !strings.Contains(path, "/") { // no import required
			continue
		}

		lastDotIndex := strings.LastIndex(path, ".")
		lastSlackIndex := strings.LastIndex(path, "/")

		scalarTypeName := path[lastDotIndex+1:]
		scalarPkgName := path[lastSlackIndex+1 : lastDotIndex]
		pkgPath := path[:lastDotIndex]

		pkg, ok := names[scalarPkgName]
		if ok && pkg != pkgPath {
			pathToParentPkg := path[:lastSlackIndex]
			parentPkgName := pathToParentPkg[strings.LastIndex(pathToParentPkg, "/")+1:]
			alias := fmt.Sprintf("%s_%s", parentPkgName, scalarPkgName)

			// modify provided input to use during codegen
			cfg.ScalarTypeMap[scalarName] = fmt.Sprintf("%s.%s", alias, scalarTypeName)

			result.Imports[pkgPath] = alias

			names[alias] = pkgPath

			continue
		}

		// modify provided input to use during codegen
		cfg.ScalarTypeMap[scalarName] = fmt.Sprintf("%s.%s", scalarPkgName, scalarTypeName)

		result.Imports[pkgPath] = ""

		names[scalarPkgName] = pkgPath
	}

	for _, o := range operations.Operations {
		responseTypes := getResponseTypes(cfg, result, o)

		result.Operations = append(result.Operations, &fezzik_ast.OperationInfo{
			Name:             o.Name,
			OperationType:    o.Operation,
			Operation:        getOperation(o),
			ResponseType:     responseTypes[0],  // root is always the first one
			ResponseSubTypes: responseTypes[1:], // root's subtypes
			Inputs:           getOperationInputs(cfg, result, o.VariableDefinitions),
			Source: fezzik_ast.Source{
				FileName: o.Position.Src.Name,
				Line:     o.Position.Line,
			},
		})
	}

	return result
}

func appendQueryDoc(doc *ast.QueryDocument, addDoc *ast.QueryDocument) {
	names := make(map[string]struct{})
	for _, opCurr := range doc.Operations {
		names[opCurr.Name] = struct{}{}
	}

	for _, opAdd := range addDoc.Operations {
		if _, found := names[opAdd.Name]; found {
			common.Check(fmt.Errorf("duplicate operation with name '%s'", opAdd.Name))
		}
	}

	doc.Operations = append(doc.Operations, addDoc.Operations...)
}

func getInputFields(cfg *config.Config, doc *fezzik_ast.Document, def *ast.Definition) []fezzik_ast.InputField {
	if def.Kind != ast.InputObject {
		panic("invalid definition")
	}

	fields := make([]fezzik_ast.InputField, len(def.Fields))
	for i := 0; i < len(def.Fields); i++ {
		fieldInfo := def.Fields[i]

		fields[i].Name = fieldInfo.Name
		fields[i].Type = getTypeInfo(fieldInfo.Type)
		fields[i].TypeName = getGoType(cfg, fieldInfo.Type)
		switch cfg.StructTagCase {
		case "snake":
			fields[i].StructTag = strcase.ToSnake(fieldInfo.Name)
		case "kebob":
			fields[i].StructTag = strcase.ToKebab(fieldInfo.Name)
		case "camel":
			fields[i].StructTag = strcase.ToLowerCamel(fieldInfo.Name)
		default:
			fields[i].StructTag = fieldInfo.Name
		}
	}
	return fields
}

func getEnumValues(def *ast.Definition) []string {
	if def.Kind != ast.Enum {
		panic("invalid definition")
	}

	values := make([]string, len(def.EnumValues))
	for i := 0; i < len(def.EnumValues); i++ {
		values[i] = def.EnumValues[i].Name
	}
	return values
}

func getOperation(operation *ast.OperationDefinition) string {
	b := common.NewStringBuilder()

	var visitSelectionSet func(ss ast.SelectionSet, indent string)
	visitSelectionSet = func(ss ast.SelectionSet, indent string) {
		if len(ss) > 0 {
			b.Write(" {")
			for i := 0; i < len(ss); i++ {
				switch fs := ss[i].(type) {
				case *ast.Field:
					if fs.Alias != fs.Name {
						b.Write("\n", indent, "\t", fs.Alias, " : ", fs.Name)
					} else {
						b.Write("\n", indent, "\t", fs.Name)
					}
					if len(fs.Arguments) > 0 {
						b.Write("(")
						for ia := 0; ia < len(fs.Arguments); ia++ {
							arg := fs.Arguments[ia]
							b.Write(arg.Name, ": ", arg.Value.String())
							if ia < len(fs.Arguments)-1 {
								b.Write(", ")
							}
						}
						b.Write(")")
					}
					visitSelectionSet(fs.SelectionSet, indent+"\t")
				default:
					common.Check(fmt.Errorf("unimplamented type %T", ss[i]))
				}

				if i == len(ss)-1 {
					b.Writeln()
				}
			}
			b.Write(indent, "}")
		}
	}

	b.Write(string(operation.Operation), " ", operation.Name)
	if len(operation.VariableDefinitions) > 0 {
		b.Write("(")
		for i := 0; i < len(operation.VariableDefinitions); i++ {
			def := operation.VariableDefinitions[i]
			b.Write("$", def.Variable, " : ", def.Type.String())
			if i < len(operation.VariableDefinitions)-1 {
				b.Write(", ")
			}
		}
		b.Write(")")
	}
	visitSelectionSet(operation.SelectionSet, "\t")

	return b.String()
}

func getResponseTypes(cfg *config.Config, doc *fezzik_ast.Document, operation *ast.OperationDefinition) []string {
	var visitSelectionSet func(ast.SelectionSet, string, string) []string
	visitSelectionSet = func(ss ast.SelectionSet, operationName, sig string) []string {
		if len(ss) == 0 {
			return nil
		}

		var allTypes []string

		b := common.NewStringBuilder()

		if sig != "" {
			b.Writeln("type ", operationName, " ", sig, " ")
		}

		b.Writeln("{")

		for _, s := range ss {
			switch fs := s.(type) {
			case *ast.Field:
				name := common.UppercaseFirstChar(fs.Alias)
				sig, bound := getFieldTypeSignature(cfg, doc, fs.Definition.Type)

				// stop if underlying selection set is empty or bound go type is found
				if len(fs.SelectionSet) == 0 || bound {
					b.Writeln(" ", name, " ", sig)

					break
				}

				var ptr = "*"
				if !strings.HasPrefix(sig, ptr) {
					ptr = ""
				}

				allTypes = append(allTypes, visitSelectionSet(fs.SelectionSet, operationName+name, strings.TrimPrefix(sig, ptr))...)

				b.Writeln(" ", name, " ", ptr+operationName+name)
			default:
				common.Check(fmt.Errorf("unimplamented type %T", s))
			}
		}

		b.Writeln("}")

		return append([]string{b.String()}, allTypes...)
	}

	return visitSelectionSet(operation.SelectionSet, operation.Name, "")
}

func getFieldTypeSignature(cfg *config.Config, doc *fezzik_ast.Document, fieldType *ast.Type) (string, bool) {
	b := common.NewStringBuilder()
	typeInfo := getTypeInfo(fieldType)

	if typeInfo.IsList {
		if typeInfo.IsListNullable {
			b.Write("*")
		}
		b.Write("[]")
	}
	if typeInfo.IsTypeNullable {
		b.Write("*")
	}

	// search among bound types
	if name, found := doc.BoundGoTypes[typeInfo.Name]; found {
		b.Write(name)
		return b.String(), true
	}

	if goType, ok := fezzik_ast.GetGoType(cfg, typeInfo.Name); ok {
		b.Write(goType)
	} else if _, found := doc.InputTypes[typeInfo.Name]; found {
		b.Write(typeInfo.Name)
	} else if _, found = doc.EnumTypes[typeInfo.Name]; found {
		b.Write(typeInfo.Name)
	} else {
		b.Write("struct")
	}
	return b.String(), false
}

func getOperationInputs(cfg *config.Config, doc *fezzik_ast.Document, vars ast.VariableDefinitionList) []fezzik_ast.OperationInput {
	result := make([]fezzik_ast.OperationInput, len(vars))
	for i := 0; i < len(vars); i++ {
		typ, _ := getFieldTypeSignature(cfg, doc, vars[i].Type)
		result[i] = fezzik_ast.OperationInput{
			Name: vars[i].Variable,
			Type: typ,
		}
	}
	return result
}

func getTypeInfo(t *ast.Type) *fezzik_ast.TypeInfo {
	if t.NamedType != "" {
		return &fezzik_ast.TypeInfo{
			Name:           t.NamedType,
			IsTypeNullable: !t.NonNull,
		}
	}

	return &fezzik_ast.TypeInfo{
		Name:           t.Elem.NamedType,
		IsList:         true,
		IsTypeNullable: !t.Elem.NonNull,
		IsListNullable: !t.NonNull,
	}
}

func getGoType(cfg *config.Config, t *ast.Type) string {
	typeName := t.Name()
	if goType, ok := fezzik_ast.GetGoType(cfg, t.Name()); ok {
		return goType
	}
	return typeName
}
