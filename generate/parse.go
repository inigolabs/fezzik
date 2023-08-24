package generate

import (
	"path/filepath"
	"regexp"
	"strings"

	"github.com/jensneuse/graphql-go-tools/pkg/federation"
	"github.com/vektah/gqlparser/v2"
	"github.com/vektah/gqlparser/v2/ast"
	"github.com/vektah/gqlparser/v2/parser"
	"github.com/vektah/gqlparser/v2/validator"

	"github.com/inigolabs/fezzik/common"
	"github.com/inigolabs/fezzik/config"
)

var reScalar = regexp.MustCompile(`scalar \w+`)

func ParseSchema(cfg *config.Config) *ast.Schema {
	// get schema files from config globs
	var schemaFiles []string
	for _, schemaGlob := range cfg.Schemas {
		files, err := filepath.Glob(schemaGlob)
		common.Check(err)
		schemaFiles = append(schemaFiles, files...)
	}

	// make sources
	var sources = make([]string, len(schemaFiles))
	for i, filename := range schemaFiles {
		sources[i] = common.FileRead(filename)
	}

	mergedSchema, err := federation.BuildBaseSchemaDocument(sources...)
	common.Check(err)

	scalars := reScalar.FindAllString(mergedSchema, -1)

	var set = map[string]struct{}{}

	for i := range scalars {
		if _, ok := set[scalars[i]]; !ok {
			set[scalars[i]] = struct{}{}

			continue
		}

		mergedSchema = strings.ReplaceAll(mergedSchema, scalars[i], "")
		mergedSchema += "\n" + scalars[i]
	}

	schema, err := gqlparser.LoadSchema(&ast.Source{Input: mergedSchema})
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
