package generate

import (
	"path/filepath"
	"strings"

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

	// make sources
	sources := make([]*ast.Source, len(schemaFiles)+1)
	sources[0] = validator.Prelude
	for i, filename := range schemaFiles {
		sources[i+1] = &ast.Source{Input: common.FileRead(filename)}
	}

	doc, err := parser.ParseSchemas(sources...)
	common.Check(err)

	// remove duplicated scalars
	defs := make([]*ast.Definition, 0, len(doc.Definitions))
	scalarsSet := make(map[string]struct{})
	for _, def := range doc.Definitions {
		if def.Kind != ast.Scalar {
			defs = append(defs, def)
			continue
		}

		// skip known scalars
		if _, ok := scalarsSet[def.Name]; ok {
			continue
		}

		defs = append(defs, def)
		scalarsSet[def.Name] = struct{}{}
	}

	doc.Definitions = defs // override definitions

	schema, err := validator.ValidateSchemaDocument(doc)
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
		source.Input = strings.Replace(source.Input, "...", "__typename ...", -1)

		doc, err := parser.ParseQuery(source)
		common.Check(err)

		result.Fragments = append(result.Fragments, doc.Fragments...)
		result.Operations = append(result.Operations, doc.Operations...)
	}

	for i := range result.Operations {
		result.Operations[i].SelectionSet = removeTypenameDuplicates(result.Operations[i].SelectionSet, result.Fragments)
	}

	err := validator.Validate(schema, result)
	common.Check(err)

	return result
}

func removeTypenameDuplicates(set []ast.Selection, fragments ast.FragmentDefinitionList) []ast.Selection {
	var newSet []ast.Selection
	var hasTypeName bool
	for i := range set {
		switch f := set[i].(type) {
		case *ast.Field:
			if f.Name == "__typename" {
				if hasTypeName {
					continue
				}

				hasTypeName = true
			}
			f.SelectionSet = removeTypenameDuplicates(f.SelectionSet, fragments)
		case *ast.InlineFragment:
			f.SelectionSet = removeTypenameDuplicates(f.SelectionSet, fragments)
		case *ast.FragmentSpread:
			fragments.ForName(f.Name).SelectionSet = removeTypenameDuplicates(fragments.ForName(f.Name).SelectionSet, fragments)
		}

		newSet = append(newSet, set[i])
	}

	return newSet
}
