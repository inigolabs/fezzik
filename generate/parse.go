package generate

import (
	"os"
	"path/filepath"

	"github.com/inigolabs/fezzik/common"
	"github.com/inigolabs/fezzik/config"

	"github.com/vektah/gqlparser/v2"
	"github.com/vektah/gqlparser/v2/ast"
	"github.com/vektah/gqlparser/v2/parser"
	"github.com/vektah/gqlparser/v2/validator"
)

func ParseSchema(cfg *config.Config) *ast.Schema {
	// get schema files from config globs
	var schemaFiles []string
	for _, schemaGlob := range cfg.Schemas {
		files, err := filepath.Glob(schemaGlob)
		common.Check(err)
		schemaFiles = append(schemaFiles, files...)
	}

	// make sources
	currDir, err := os.Getwd()
	common.Check(err)
	sources := make([]*ast.Source, len(schemaFiles))
	for i, filename := range schemaFiles {
		sources[i] = &ast.Source{
			Name:  filepath.Join(currDir, filename),
			Input: common.FileRead(filename),
		}
	}
	schema, err := gqlparser.LoadSchema(sources...)
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
	currDir, err := os.Getwd()
	common.Check(err)
	sources := make([]*ast.Source, len(operationFiles))
	for i, filename := range operationFiles {
		sources[i] = &ast.Source{
			Name:  filepath.Join(currDir, filename),
			Input: common.FileRead(filename),
		}
	}

	result := &ast.QueryDocument{}
	for _, source := range sources {
		doc, err := parser.ParseQuery(source)
		common.Check(err)
		appendQueryDoc(result, doc)
	}

	err = validator.Validate(schema, result)
	common.Check(err)

	return result
}
