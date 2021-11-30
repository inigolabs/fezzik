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

func ParseInputs(cfg *config.Config, schema *ast.Document, doc *fezzik_ast.Document) error {
	walker := astvisitor.NewWalker(48)
	visitor := &inputVisitor{
		Walker: &walker,
		config: cfg,
		schema: schema,
		doc:    doc,
	}
	walker.RegisterEnterInputObjectTypeDefinitionVisitor(visitor)
	report := &operationreport.Report{}
	visitor.Walker.Walk(visitor.schema, visitor.schema, report)
	if report.HasErrors() {
		log.Error().Str("errors", report.Error()).Msg("")
		return errors.New(report.Error())
	}
	return nil
}

type inputVisitor struct {
	*astvisitor.Walker

	config *config.Config
	schema *ast.Document

	doc *fezzik_ast.Document
}

func (v *inputVisitor) EnterInputObjectTypeDefinition(ref int) {
	name := v.schema.InputObjectTypeDefinitionNameString(ref)
	inputType := fezzik_ast.InputType{
		Name: name,
	}

	for _, fieldRef := range v.schema.InputObjectTypeDefinitions[ref].InputFieldsDefinition.Refs {
		fieldName := v.schema.InputValueDefinitionNameString(fieldRef)
		typeRef := v.schema.InputValueDefinitionType(fieldRef)
		typeInfo := fezzik_ast.GetTypeInfo(typeRef, v.schema)

		inputField := fezzik_ast.InputField{
			Name: fieldName,
			Type: typeInfo,
		}
		if goType, ok := fezzik_ast.GetGoType(v.config, typeInfo.Name); ok {
			inputField.TypeName = goType
		} else {
			inputField.TypeName = typeInfo.Name
		}
		inputType.Fields = append(inputType.Fields, inputField)
	}

	v.doc.InputTypes = append(v.doc.InputTypes, inputType)
}
