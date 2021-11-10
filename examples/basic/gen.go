package main

import (
	"os"

	"github.com/inigolabs/fezzik"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

//go:generate go run gen.go

func init() {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	log.Logger = log.With().Caller().Logger().Output(zerolog.ConsoleWriter{Out: os.Stderr})
}

func main() {
	cfg := &fezzik.Config{
		Schemas:    []string{"schemas/*.graphql"},
		Operations: []string{"operations/*.graphql"},

		PackageName: "basic",
		PackageDir:  "gen",
		ScalarTypeMap: map[string]string{
			"Time": "time.Time",
		},
	}
	fezzik.Generate(cfg)
}
