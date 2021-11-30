package parse

import (
	"errors"

	"github.com/inigolabs/fezzik/config"
	"github.com/inigolabs/fezzik/fezzik_ast"
	"github.com/jensneuse/graphql-go-tools/pkg/ast"
	"github.com/jensneuse/graphql-go-tools/pkg/astvisitor"
	"github.com/jensneuse/graphql-go-tools/pkg/operationreport"
	"github.com/rs/zerolog/log"
)

func ParseEnums(cfg *config.Config, schema *ast.Document, doc *fezzik_ast.Document) error {
	walker := astvisitor.NewWalker(48)
	visitor := &enumVisitor{
		Walker: &walker,
		config: cfg,
		schema: schema,
		doc:    doc,
	}
	walker.RegisterEnterEnumTypeDefinitionVisitor(visitor)
	report := &operationreport.Report{}
	visitor.Walker.Walk(visitor.schema, visitor.schema, report)
	if report.HasErrors() {
		log.Error().Str("errors", report.Error()).Msg("")
		return errors.New(report.Error())
	}
	return nil
}

type enumVisitor struct {
	*astvisitor.Walker

	config *config.Config
	schema *ast.Document

	doc *fezzik_ast.Document
}

func (v *enumVisitor) EnterEnumTypeDefinition(ref int) {
	enumType := fezzik_ast.EnumType{
		Name: v.schema.EnumTypeDefinitionNameString(ref),
	}
	refs := v.schema.EnumTypeDefinitions[ref].EnumValuesDefinition.Refs
	for _, valueRef := range refs {
		enumType.Values = append(enumType.Values,
			v.schema.EnumValueDefinitionNameString(valueRef))
	}

	v.doc.EnumTypes[enumType.Name] = enumType
}
