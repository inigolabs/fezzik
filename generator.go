package fezzik

import (
	"os"
	"path/filepath"

	"github.com/inigolabs/fezzik/config"
	"github.com/inigolabs/fezzik/fezzik_ast"
	"github.com/inigolabs/fezzik/parse"
	"github.com/inigolabs/fezzik/render"
)

func Generate(cfg *config.Config) {

	doc := fezzik_ast.NewDocument()

	// parse
	schema := parse.SchemaDoc(cfg)
	check(parse.ParseEnums(cfg, schema, doc))
	check(parse.ParseInputs(cfg, schema, doc))
	check(parse.ParseOperations(cfg, schema, doc))

	// render
	doc.PackageName = cfg.PackageName
	genPath := filepath.Join(cfg.PackageDir, cfg.PackageName)
	check(os.MkdirAll(genPath, os.ModePerm))

	writer, err := os.Create(filepath.Join(genPath, "inputs.go"))
	check(err)
	render.Render(cfg, "inputs.go.tmpl", doc, writer)

	writer, err = os.Create(filepath.Join(genPath, "operations.go"))
	check(err)
	render.Render(cfg, "operations.go.tmpl", doc, writer)

	writer, err = os.Create(filepath.Join(genPath, "client.go"))
	check(err)
	render.Render(cfg, "client.go.tmpl", doc, writer)

	if cfg.GenerateMocks {
		render.RenderMocks(cfg)
	}
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
