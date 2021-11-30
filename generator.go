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
	"github.com/inigolabs/fezzik/config"
	"github.com/inigolabs/fezzik/fezzik_ast"
	"github.com/inigolabs/fezzik/parse"
	"github.com/inigolabs/fezzik/render"
	"github.com/jensneuse/graphql-go-tools/pkg/ast"
	"github.com/jensneuse/graphql-go-tools/pkg/astnormalization"
	"github.com/jensneuse/graphql-go-tools/pkg/astparser"
	"github.com/jensneuse/graphql-go-tools/pkg/astvalidation"
	"github.com/jensneuse/graphql-go-tools/pkg/operationreport"
	"github.com/rs/zerolog/log"
	"golang.org/x/tools/imports"

	"github.com/iancoleman/strcase"
	mockery "github.com/vektra/mockery/v2/pkg"
	mockery_config "github.com/vektra/mockery/v2/pkg/config"
)

func Generate(cfg *config.Config) {

	schema := schemaDoc(cfg)
	doc := fezzik_ast.NewDocument()

	// parse
	check(parse.ParseEnums(cfg, schema, doc))
	check(parse.ParseInputs(cfg, schema, doc))

	var operationFiles []string
	for _, operationGlob := range cfg.Operations {
		files, err := filepath.Glob(operationGlob)
		check(err)
		operationFiles = append(operationFiles, files...)
	}
	operationParser := parse.NewOperationParser(cfg, schema, doc)
	for _, operationFile := range operationFiles {
		operationStr := common.FileRead(operationFile)
		check(operationParser.ParseOperation(operationStr))
	}

	// render
	genPath := filepath.Join(cfg.PackageDir, cfg.PackageName)
	check(os.MkdirAll(genPath, os.ModePerm))
	doc.PackageName = cfg.PackageName
	writer, err := os.Create(filepath.Join(genPath, "inputs.go"))
	check(err)
	templateStr := strings.ReplaceAll(render.InputsTemplate, "~~", "`")
	generate(cfg, templateStr, doc, writer)

	writer, err = os.Create(filepath.Join(genPath, "operations.go"))
	check(err)
	templateStr = strings.ReplaceAll(render.OperationsTemplate, "~~", "`")
	generate(cfg, templateStr, doc, writer)

	writer, err = os.Create(filepath.Join(genPath, "client.go"))
	check(err)
	templateStr = strings.ReplaceAll(render.ClientTemplate, "~~", "`")
	generate(cfg, templateStr, doc, writer)

	if cfg.GenerateMocks {
		generateMocks(cfg)
	}
}

func schemaDoc(cfg *config.Config) *ast.Document {
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

func generate(cfg *config.Config, templateStr string, info interface{}, writer io.Writer) {
	templateFuncs := template.FuncMap{
		"pascal": strcase.ToCamel,
	}
	template := template.Must(template.New("tmpl").Funcs(templateFuncs).Parse(templateStr))
	var out bytes.Buffer
	template.Execute(&out, info)
	formatted, err := imports.Process("", out.Bytes(), nil)
	if err != nil {
		if cfg.Debug {
			log.Error().Err(err).Msg("error formatting generated code")
			formatted = out.Bytes()
		} else {
			panic("error formatting generated code")
		}
	}
	_, err = writer.Write(formatted)
	check(err)
}

func generateMocks(cfg *config.Config) {
	mcfg := mockery_config.Config{
		Name:           "Client",
		Case:           "underscore",
		Dir:            ".",
		LogLevel:       "error",
		Outpkg:         "mocks",
		Output:         filepath.Join(cfg.PackageDir, cfg.PackageName, "mocks"),
		UnrollVariadic: true,
	}

	osp := &mockery.FileOutputStreamProvider{
		Config:                    mcfg,
		BaseDir:                   mcfg.Output,
		InPackage:                 mcfg.InPackage,
		TestOnly:                  mcfg.TestOnly,
		Case:                      mcfg.Case,
		KeepTree:                  mcfg.KeepTree,
		KeepTreeOriginalDirectory: mcfg.Dir,
		FileName:                  mcfg.FileName,
	}

	visitor := &mockery.GeneratorVisitor{
		Config:            mcfg,
		InPackage:         mcfg.InPackage,
		Note:              mcfg.Note,
		Boilerplate:       "",
		Osp:               osp,
		PackageName:       mcfg.Outpkg,
		PackageNamePrefix: mcfg.Packageprefix,
		StructName:        mcfg.StructName,
	}

	walker := mockery.Walker{
		Config:    mcfg,
		BaseDir:   mcfg.Dir,
		Recursive: true,
		Filter:    regexp.MustCompile("Client"),
		LimitOne:  true,
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
