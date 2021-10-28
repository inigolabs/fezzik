package fezzik

import (
	"bytes"
	"context"
	"errors"
	"html/template"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/inigolabs/fezzik/common"
	"github.com/jensneuse/graphql-go-tools/pkg/ast"
	"github.com/jensneuse/graphql-go-tools/pkg/astnormalization"
	"github.com/jensneuse/graphql-go-tools/pkg/astparser"
	"github.com/jensneuse/graphql-go-tools/pkg/astvalidation"
	"github.com/jensneuse/graphql-go-tools/pkg/operationreport"
	"github.com/rs/zerolog/log"
	"golang.org/x/tools/imports"

	mockery "github.com/vektra/mockery/v2/pkg"
	mockery_config "github.com/vektra/mockery/v2/pkg/config"
)

func Generate(cfg *Config) {
	schemaFiles, err := filepath.Glob(cfg.Schemas)
	check(err)

	// parse schema
	schemaBuilder := new(strings.Builder)
	schemaBuilder.WriteString("scalar ID\n")
	schemaBuilder.WriteString("scalar Int\n")
	schemaBuilder.WriteString("scalar String\n")
	schemaBuilder.WriteString("scalar Boolean\n")
	schemaBuilder.WriteString("scalar Float\n")
	for _, schemaFileName := range schemaFiles {
		_, err := io.Copy(schemaBuilder, common.FileReader(schemaFileName))
		check(err)
	}
	schema, err := parseSchema(schemaBuilder.String())
	check(err)

	if cfg.PackageName == "mocks" {
		panic("PackageName cannot be 'mocks'")
	}
	genPath := filepath.Join(cfg.PackageDir, cfg.PackageName)
	err = os.MkdirAll(genPath, os.ModePerm)
	check(err)

	// walk definition
	inputVisitor := NewInputVisitor(schema, cfg)
	inputVisitor.Walk()
	writer, err := os.Create(filepath.Join(genPath, "inputs.go"))
	check(err)
	templateStr := strings.ReplaceAll(inputsTemplate, "~~", "`")
	generate(templateStr, inputVisitor.info, writer)

	// walk operations
	visitor := NewVisitor(cfg, schema, inputVisitor.info)
	operationFiles, err := filepath.Glob(cfg.Operations)
	check(err)
	for _, operationFile := range operationFiles {
		operationStr := common.FileRead(operationFile)
		operation, err := parseOperation(operationStr, schema)
		check(err)
		err = visitor.AddInfo(operationStr, operation)
		check(err)
	}

	writer, err = os.Create(filepath.Join(genPath, "operations.go"))
	check(err)
	templateStr = strings.ReplaceAll(operationsTemplate, "~~", "`")
	generate(templateStr, visitor.info, writer)

	writer, err = os.Create(filepath.Join(genPath, "client.go"))
	check(err)
	templateStr = strings.ReplaceAll(clientTemplate, "~~", "`")
	generate(templateStr, visitor.info, writer)
	generateMocks()
}

func parseSchema(schemaString string) (*ast.Document, error) {
	parser := astparser.NewParser()
	doc := ast.NewDocument()
	doc.Input.ResetInputBytes([]byte(schemaString))
	report := operationreport.Report{}
	parser.Parse(doc, &report)
	if report.HasErrors() {
		return nil, errors.New(report.Error())
	}

	astnormalization.NormalizeDefinition(doc, &report)
	if report.HasErrors() {
		return nil, errors.New(report.Error())
	}

	validator := astvalidation.DefaultDefinitionValidator()
	report = operationreport.Report{}
	result := validator.Validate(doc, &report)
	if result != astvalidation.Valid {
		return nil, errors.New(report.Error())
	}
	return doc, nil
}

func parseOperation(operationString string, schema *ast.Document) (*ast.Document, error) {
	parser := astparser.NewParser()
	doc := ast.NewDocument()
	doc.Input.ResetInputBytes([]byte(operationString))
	report := operationreport.Report{}
	parser.Parse(doc, &report)
	if report.HasErrors() {
		return nil, errors.New(report.Error())
	}

	validator := astvalidation.DefaultOperationValidator()
	report = operationreport.Report{}
	result := validator.Validate(doc, schema, &report)
	if result != astvalidation.Valid {
		return nil, errors.New(report.Error())
	}
	return doc, nil
}

func generate(templateStr string, info interface{}, writer io.Writer) {
	templateFuncs := template.FuncMap{
		"capitalize": strings.Title,
	}
	template := template.Must(template.New("tmpl").Funcs(templateFuncs).Parse(templateStr))
	var out bytes.Buffer
	template.Execute(&out, info)
	formatted, err := imports.Process("", out.Bytes(), nil)
	if err != nil {
		formatted = out.Bytes()
		log.Error().Err(err).Msg("formatting error")
	}
	_, err = writer.Write(formatted)
	check(err)
}

func generateMocks() {
	cfg := mockery_config.Config{
		All:            true,
		Case:           "underscore",
		Dir:            ".",
		LogLevel:       "error",
		Outpkg:         "mocks",
		Output:         "gen/mocks",
		UnrollVariadic: true,
	}

	osp := &mockery.FileOutputStreamProvider{
		Config:                    cfg,
		BaseDir:                   cfg.Output,
		InPackage:                 cfg.InPackage,
		TestOnly:                  cfg.TestOnly,
		Case:                      cfg.Case,
		KeepTree:                  cfg.KeepTree,
		KeepTreeOriginalDirectory: cfg.Dir,
		FileName:                  cfg.FileName,
	}

	visitor := &mockery.GeneratorVisitor{
		Config:            cfg,
		InPackage:         cfg.InPackage,
		Note:              cfg.Note,
		Boilerplate:       "",
		Osp:               osp,
		PackageName:       cfg.Outpkg,
		PackageNamePrefix: cfg.Packageprefix,
		StructName:        cfg.StructName,
	}

	walker := mockery.Walker{
		Config:    cfg,
		BaseDir:   cfg.Dir,
		Recursive: true,
		Filter:    regexp.MustCompile(".*"),
		LimitOne:  false,
	}

	ctx := context.Background()
	generated := walker.Walk(ctx, visitor)
	if !generated {
		panic("mocks not generated")
	}
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
