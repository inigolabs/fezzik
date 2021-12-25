package main

import (
	"os"

	"github.com/alexflint/go-arg"
	"github.com/ejoffe/rake"
	"github.com/inigolabs/fezzik/config"
	"github.com/inigolabs/fezzik/generate"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func init() {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	log.Logger = log.With().Caller().Logger().Output(zerolog.ConsoleWriter{Out: os.Stderr})
}

func main() {
	var args struct {
		Config string `arg:"required" help:"fezzik yaml config file path"`
	}
	arg.MustParse(&args)

	var cfg config.Config
	rake.LoadSources(&cfg,
		rake.DefaultSource(),
		rake.YamlMustFileSource(args.Config),
	)

	if cfg.Debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}
	generate.Generate(&cfg)
}
