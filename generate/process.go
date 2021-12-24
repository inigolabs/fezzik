package generate

import (
	"fmt"

	"github.com/inigolabs/fezzik/common"
	"github.com/inigolabs/fezzik/config"
	"github.com/inigolabs/fezzik/fezzik_ast"

	"github.com/vektah/gqlparser/v2/ast"
)

func Process(cfg *config.Config, schema *ast.Schema, operations *ast.QueryDocument) *fezzik_ast.Document {
	result := &fezzik_ast.Document{
		InputTypes: make(map[string]fezzik_ast.InputType),
		EnumTypes:  make(map[string]fezzik_ast.EnumType),
	}

	for _, t := range schema.Types {
		if t.Kind == ast.Enum {
			result.EnumTypes[t.Name] = fezzik_ast.EnumType{
				Name:   t.Name,
				Values: getEnumValues(t),
			}
		}
	}

	for _, t := range schema.Types {
		if t.Kind == ast.InputObject {
			result.InputTypes[t.Name] = fezzik_ast.InputType{
				Name:   t.Name,
				Fields: getInputFields(cfg, result, t),
			}
		}
	}

	for _, o := range operations.Operations {
		result.Operations = append(result.Operations, &fezzik_ast.OperationInfo{
			Name:         o.Name,
			Operation:    getOperation(o),
			ResponseType: getResponseType(cfg, result, o),
			Inputs:       getOperationInputs(cfg, result, o.VariableDefinitions),
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
					b.Write("\n", indent, "   ", fs.Name)
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
					visitSelectionSet(fs.SelectionSet, indent+"   ")
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
	visitSelectionSet(operation.SelectionSet, "")

	return b.String()
}

func getResponseType(cfg *config.Config, doc *fezzik_ast.Document, operation *ast.OperationDefinition) string {
	b := common.NewStringBuilder()

	var visitSelectionSet func(ss ast.SelectionSet, indent string)
	visitSelectionSet = func(ss ast.SelectionSet, indent string) {
		if len(ss) > 0 {
			b.Writeln(indent, "{")
			for _, s := range ss {
				switch fs := s.(type) {
				case *ast.Field:
					name := common.UppercaseFirstChar(fs.Name)
					sig := getFieldTypeSignature(cfg, doc, fs.Definition.Type)
					b.Writeln(indent, name, " ", sig)
					visitSelectionSet(fs.SelectionSet, indent+"  ")
				default:
					common.Check(fmt.Errorf("unimplamented type %T", s))
				}

			}
			b.Writeln(indent, "}")
		}
	}

	visitSelectionSet(operation.SelectionSet, "")

	return b.String()
}

func getFieldTypeSignature(cfg *config.Config, doc *fezzik_ast.Document, fieldType *ast.Type) string {
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
	if goType, ok := fezzik_ast.GetGoType(cfg, typeInfo.Name); ok {
		b.Write(goType)
	} else if _, found := doc.InputTypes[typeInfo.Name]; found {
		b.Write(typeInfo.Name)
	} else if _, found := doc.EnumTypes[typeInfo.Name]; found {
		b.Write(typeInfo.Name)
	} else {
		b.Write("struct")
	}
	return b.String()
}

func getOperationInputs(cfg *config.Config, doc *fezzik_ast.Document, vars ast.VariableDefinitionList) []fezzik_ast.OperationInput {
	result := make([]fezzik_ast.OperationInput, len(vars))
	for i := 0; i < len(vars); i++ {
		result[i] = fezzik_ast.OperationInput{
			Name: vars[i].Variable,
			Type: getFieldTypeSignature(cfg, doc, vars[i].Type),
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
