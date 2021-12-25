package render

import (
	"bytes"
	"context"
	"embed"
	"html/template"
	"io"
	"path/filepath"
	"regexp"

	"github.com/iancoleman/strcase"
	"github.com/inigolabs/fezzik/config"
	"github.com/inigolabs/fezzik/fezzik_ast"
	"github.com/rs/zerolog/log"
	"golang.org/x/tools/imports"

	mockery "github.com/vektra/mockery/v2/pkg"
	mockery_config "github.com/vektra/mockery/v2/pkg/config"
)

var (
	//go:embed inputs.go.tmpl
	//go:embed operations.go.tmpl
	//go:embed client.go.tmpl
	templateFiles embed.FS
)

func Render(cfg *config.Config, templateName string, doc *fezzik_ast.Document, writer io.Writer) {
	templateFuncs := template.FuncMap{
		"pascal": strcase.ToCamel,
		"camel":  strcase.ToLowerCamel,
	}
	template := template.Must(template.New(templateName).Funcs(templateFuncs).
		ParseFS(templateFiles, templateName))
	var out bytes.Buffer
	err := template.Execute(&out, doc)
	check(err)
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

func RenderMocks(cfg *config.Config) {
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
