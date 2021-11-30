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

		PackageName: "github",
		PackageDir:  "gen",
		ScalarBuiltIn: []string{
			"Int", "ID", "String", "Boolean", "Float",
		},
		ScalarTypeMap: map[string]string{
			"DateTime":    "time.Time",
			"URI":         "string",
			"GitObjectID": "string",
			"GitRefname":  "string",
		},
		GenerateMocks: true,
	}
	fezzik.Generate(cfg)
}
