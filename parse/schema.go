package parse

import (
	"errors"
	"io"
	"path/filepath"
	"strings"

	"github.com/inigolabs/fezzik/common"
	"github.com/inigolabs/fezzik/config"
	"github.com/jensneuse/graphql-go-tools/pkg/ast"
	"github.com/jensneuse/graphql-go-tools/pkg/astnormalization"
	"github.com/jensneuse/graphql-go-tools/pkg/astparser"
	"github.com/jensneuse/graphql-go-tools/pkg/astvalidation"
	"github.com/jensneuse/graphql-go-tools/pkg/operationreport"
)

func SchemaDoc(cfg *config.Config) *ast.Document {
	var schemaFiles []string
	for _, schemaGlob := range cfg.Schemas {
		files, err := filepath.Glob(schemaGlob)
		check(err)
		schemaFiles = append(schemaFiles, files...)
	}
	// parse schema
	schemaBuilder := new(strings.Builder)
	for _, scalar := range cfg.ScalarBuiltIn {
		schemaBuilder.WriteString("scalar " + scalar + "\n")
	}

	for _, schemaFileName := range schemaFiles {
		_, err := io.Copy(schemaBuilder, common.FileReader(schemaFileName))
		check(err)
	}

	parser := astparser.NewParser()
	doc := ast.NewDocument()
	doc.Input.ResetInputBytes([]byte(schemaBuilder.String()))
	report := operationreport.Report{}
	parser.Parse(doc, &report)
	if report.HasErrors() {
		panic(errors.New(report.Error()))
	}

	astnormalization.NormalizeDefinition(doc, &report)
	if report.HasErrors() {
		panic(errors.New(report.Error()))
	}

	validator := astvalidation.DefaultDefinitionValidator()
	report = operationreport.Report{}
	result := validator.Validate(doc, &report)
	if result != astvalidation.Valid {
		panic(errors.New(report.Error()))
	}
	return doc
}
