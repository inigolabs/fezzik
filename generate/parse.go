package generate

import (
	"path/filepath"
	"strings"

	"github.com/vektah/gqlparser/v2"
	"github.com/vektah/gqlparser/v2/ast"
	"github.com/vektah/gqlparser/v2/parser"
	"github.com/vektah/gqlparser/v2/validator"

	"github.com/inigolabs/fezzik/common"
	"github.com/inigolabs/fezzik/config"
)

func ParseSchema(cfg *config.Config) *ast.Schema {
	// get schema files from config globs
	var schemaFiles []string
	for _, schemaGlob := range cfg.Schemas {
		files, err := filepath.Glob(schemaGlob)
		common.Check(err)
		schemaFiles = append(schemaFiles, files...)
	}

	// combine schema files
	b := strings.Builder{}
	for _, filename := range schemaFiles {
		b.WriteRune('\n')
		b.WriteString(common.FileRead(filename))
	}

	schema, err := gqlparser.LoadSchema(&ast.Source{Input: b.String()})
	common.Check(err)
	return schema
}

func ParseOperations(cfg *config.Config, schema *ast.Schema) *ast.QueryDocument {
	var operationFiles []string
	for _, operationGlob := range cfg.Operations {
		files, err := filepath.Glob(operationGlob)
		common.Check(err)
		operationFiles = append(operationFiles, files...)
	}

	// make sources
	sources := make([]*ast.Source, len(operationFiles))
	for i, filename := range operationFiles {
		sources[i] = &ast.Source{
			Name:  common.RootPath(filename),
			Input: common.FileRead(filename),
		}
	}

	result := &ast.QueryDocument{}
	for _, source := range sources {
		doc, err := parser.ParseQuery(source)
		common.Check(err)

		result.Fragments = append(result.Fragments, doc.Fragments...)
		result.Operations = append(result.Operations, doc.Operations...)
	}

	err := validator.Validate(schema, result)
	common.Check(err)

	return result
}
