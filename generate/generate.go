package generate

import (
	"os"
	"path/filepath"

	"github.com/inigolabs/fezzik/common"
	"github.com/inigolabs/fezzik/config"
	"github.com/inigolabs/fezzik/render"
)

func Generate(cfg *config.Config) {
	if cfg.Debug {
		common.PrefixPretty("config", cfg)
	}

	// parse
	schema := ParseSchema(cfg)
	operations := ParseOperations(cfg, schema)

	// process
	doc := Process(cfg, schema, operations)
	doc.PackageName = cfg.PackageName

	if cfg.Debug {
		common.PrefixPretty("doc", doc)
	}

	// render
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

	writer, err = os.Create(filepath.Join(genPath, "subscription_client.go"))
	check(err)
	render.Render(cfg, "subscription_client.go.tmpl", doc, writer)

	if cfg.GenerateMocks {
		render.RenderMocks(cfg)
	}
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
