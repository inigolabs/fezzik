package parse

import (
	"github.com/inigolabs/fezzik/config"
	"github.com/inigolabs/fezzik/fezzik_ast"
	"github.com/jensneuse/graphql-go-tools/pkg/ast"
)

func ParseEnums(cfg *config.Config, schema *ast.Document, doc *fezzik_ast.Document) error {
	parser := &enumParser{
		config: cfg,
		schema: schema,
		doc:    doc,
	}
	for ref := range schema.EnumTypeDefinitions {
		parser.ParseEnumTypeDefinition(ref)
	}
	return nil
}

type enumParser struct {
	config *config.Config
	schema *ast.Document

	doc *fezzik_ast.Document
}

func (p *enumParser) ParseEnumTypeDefinition(ref int) {
	enumType := fezzik_ast.EnumType{
		Name: p.schema.EnumTypeDefinitionNameString(ref),
	}
	refs := p.schema.EnumTypeDefinitions[ref].EnumValuesDefinition.Refs
	for _, valueRef := range refs {
		enumType.Values = append(enumType.Values,
			p.schema.EnumValueDefinitionNameString(valueRef))
	}

	p.doc.EnumTypes[enumType.Name] = enumType
}
