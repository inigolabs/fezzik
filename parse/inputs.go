package parse

import (
	"fmt"

	"github.com/inigolabs/fezzik/config"
	"github.com/inigolabs/fezzik/fezzik_ast"
	"github.com/jensneuse/graphql-go-tools/pkg/ast"
)

func ParseInputs(cfg *config.Config, schema *ast.Document, doc *fezzik_ast.Document) error {
	parser := &inputParser{
		config: cfg,
		schema: schema,
		doc:    doc,
	}
	for ref := range schema.InputObjectTypeDefinitions {
		parser.FirstPassInputObjectTypeDefinition(ref)
	}
	for ref := range schema.InputObjectTypeDefinitions {
		parser.SecondPassInputObjectTypeDefinition(ref)
	}
	return nil
}

type inputParser struct {
	config *config.Config
	schema *ast.Document

	doc *fezzik_ast.Document
}

func (p *inputParser) FirstPassInputObjectTypeDefinition(ref int) {
	name := p.schema.InputObjectTypeDefinitionNameString(ref)
	p.doc.InputTypes[name] = fezzik_ast.InputType{
		Name: name,
	}
}

func (p *inputParser) SecondPassInputObjectTypeDefinition(ref int) {
	name := p.schema.InputObjectTypeDefinitionNameString(ref)
	inputType := fezzik_ast.InputType{
		Name: name,
	}

	for _, fieldRef := range p.schema.InputObjectTypeDefinitions[ref].InputFieldsDefinition.Refs {
		fieldName := p.schema.InputValueDefinitionNameString(fieldRef)
		typeRef := p.schema.InputValueDefinitionType(fieldRef)
		typeInfo := fezzik_ast.GetTypeInfo(typeRef, p.schema)

		inputField := fezzik_ast.InputField{
			Name: fieldName,
			Type: typeInfo,
		}

		if goType, ok := fezzik_ast.GetGoType(p.config, typeInfo.Name); ok {
			inputField.TypeName = goType
		} else if _, found := p.doc.EnumTypes[typeInfo.Name]; found {
			inputField.TypeName = typeInfo.Name
		} else if _, found := p.doc.InputTypes[typeInfo.Name]; found {
			inputField.TypeName = typeInfo.Name
		} else {
			panic(fmt.Errorf("unknown type '%s' not found in config.ScalarTypeMap ", typeInfo.Name))
		}
		inputType.Fields = append(inputType.Fields, inputField)
	}
	p.doc.InputTypes[inputType.Name] = inputType
}
