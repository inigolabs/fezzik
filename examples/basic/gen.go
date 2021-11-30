package main

import (
	"os"

	"github.com/inigolabs/fezzik"
	"github.com/inigolabs/fezzik/config"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

//go:generate go run gen.go

func init() {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	log.Logger = log.With().Caller().Logger().Output(zerolog.ConsoleWriter{Out: os.Stderr})
}

func main() {
	cfg := &config.Config{
		Schemas:    []string{"schemas/*.graphql"},
		Operations: []string{"operations/*.graphql"},

		PackageName: "basic",
		PackageDir:  "gen",
		ScalarBuiltIn: []string{
			"Int", "ID", "String", "Boolean", "Float",
		},
		ScalarTypeMap: map[string]string{
			"Time": "time.Time",
		},
		GenerateMocks: true,
		Debug:         true,
	}
	fezzik.Generate(cfg)
}
