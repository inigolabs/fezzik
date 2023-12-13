package generate

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/inigolabs/fezzik/common"
	"github.com/inigolabs/fezzik/config"
	"github.com/inigolabs/fezzik/fezzik_ast"
	"github.com/vektah/gqlparser/v2/ast"
	"github.com/vektah/gqlparser/v2/formatter"
	"golang.org/x/tools/go/packages"
)

var mode = packages.NeedName | packages.NeedTypes | packages.NeedModule | packages.NeedTypesInfo

func Process(cfg *config.Config, schema *ast.Schema, operations *ast.QueryDocument) *fezzik_ast.Document {
	result := &fezzik_ast.Document{
		InputTypes:   make(map[string]fezzik_ast.InputType),
		BoundGoTypes: make(map[string]string),
		EnumTypes:    make(map[string]fezzik_ast.EnumType),
		Imports:      make(map[string]string),

		// helper to keep track of already visited fragments, so go-struct is generated only once and re-used
		VisitedFragments: map[string]struct{}{},
	}

	var pkgs []*packages.Package
	var err error
	if len(cfg.Autobind) > 0 {
		pkgs, err = packages.Load(&packages.Config{Mode: mode}, cfg.Autobind...)
		check(err)
	}

	for _, t := range schema.Types {
		if t.Kind == ast.Enum || t.Kind == ast.InputObject || t.Kind == ast.Object {
			for _, pkg := range pkgs {
				if foundType := pkg.Types.Scope().Lookup(t.Name); foundType != nil {
					result.BoundGoTypes[t.Name] = foundType.Pkg().Name() + "." + foundType.Name()
					result.Imports[foundType.Pkg().Path()] = ""
				}
			}
		}

		// add all enums
		if t.Kind == ast.Enum {
			if _, ok := result.BoundGoTypes[t.Name]; ok {
				continue
			}

			result.EnumTypes[t.Name] = fezzik_ast.EnumType{
				Name:   t.Name,
				Values: getEnumValues(t),
			}
		}
	}

	// generate only used inputs and enums
	for _, op := range operations.Operations {
		for _, v := range op.VariableDefinitions {
			def, ok := schema.Types[v.Type.Name()]
			if !ok {
				panic(fmt.Sprintf("type %s not found", v.Type.Name()))
			}

			if def.Kind == ast.InputObject {
				addInputs(def, cfg, schema, result)
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
		responseTypes := getResponseTypes(cfg, schema, result, o)

		result.Operations = append(result.Operations, &fezzik_ast.OperationInfo{
			Name:             o.Name,
			OperationType:    o.Operation,
			Operation:        getOperation(o, operations.Fragments),
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

func addInputs(def *ast.Definition, cfg *config.Config, schema *ast.Schema, result *fezzik_ast.Document) {
	if _, ok := result.BoundGoTypes[def.Name]; ok {
		return
	}

	t, ok := schema.Types[def.Name]
	if !ok {
		panic(fmt.Sprintf("type %s not found", def.Name))
	}

	result.InputTypes[def.Name] = fezzik_ast.InputType{
		Name:   def.Name,
		Fields: getInputFields(cfg, result, t),
	}

	for _, f := range t.Fields {
		fdef, ok := schema.Types[f.Type.Name()]
		if !ok {
			panic(fmt.Sprintf("type %s not found", f.Type.Name()))
		}

		if fdef.Kind == ast.InputObject {
			addInputs(fdef, cfg, schema, result)
		}
	}
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
		fields[i].StructTag = fieldInfo.Name
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

func getOperation(operation *ast.OperationDefinition, fragments ast.FragmentDefinitionList) string {
	b := bytes.Buffer{}

	used := ast.FragmentDefinitionList{}
	var findUsedFragments func(ast.SelectionSet)
	findUsedFragments = func(selectionSet ast.SelectionSet) {
		for _, selection := range selectionSet {
			switch s := selection.(type) {
			case *ast.Field:
				findUsedFragments(s.SelectionSet)
			case *ast.FragmentSpread:
				used = append(used, fragments.ForName(s.Name))
			default:
				common.Check(fmt.Errorf("unimplamented type %T", s))
			}
		}
	}

	findUsedFragments(operation.SelectionSet)

	// print operation
	formatter.NewFormatter(&b).FormatQueryDocument(&ast.QueryDocument{
		Operations: ast.OperationList{operation},
		Fragments:  used,
	})

	return b.String()
}

func getResponseTypes(cfg *config.Config, schema *ast.Schema, doc *fezzik_ast.Document, operation *ast.OperationDefinition) []string {
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
				kind := schema.Types[fs.Definition.Type.Name()].Kind
				if kind == ast.Union || kind == ast.Interface {
					panic("union and interface fields are not supported")
				}

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
			case *ast.FragmentSpread:
				if _, found := doc.VisitedFragments[fs.Name]; !found {
					doc.VisitedFragments[fs.Name] = struct{}{}
					allTypes = append(allTypes, visitSelectionSet(fs.Definition.SelectionSet, fs.Name, "struct")...)
				}

				b.Writeln(fmt.Sprintf("*%s", fs.Name))
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
