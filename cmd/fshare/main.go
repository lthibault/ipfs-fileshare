package main

import (
	"os"

	"github.com/urfave/cli/v2"

	log "github.com/lthibault/log/pkg"

	"github.com/lthibault/ipfs-fileshare/internal/cmd/repo"
	"github.com/lthibault/ipfs-fileshare/internal/cmd/start"
)

func main() {
	app := &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "logfmt",
				Aliases: []string{"f"},
				Usage:   "text, json, none",
				Value:   "text",
				EnvVars: []string{"CASM_LOGFMT"},
			},
			&cli.StringFlag{
				Name:    "loglvl",
				Usage:   "trace, debug, info, warn, error, fatal",
				Value:   "info",
				EnvVars: []string{"CASM_LOGLVL"},
			},

			/************************
			*	undocumented flags	*
			*************************/
			&cli.BoolFlag{
				Name:    "prettyprint",
				Aliases: []string{"pp"},
				Usage:   "pretty-print JSON output",
				Hidden:  true,
			},
		},
		Commands: []*cli.Command{{
			Name:   "start",
			Usage:  "start a file sharing daemon",
			Flags:  start.Flags(),
			Action: start.Run(),
		}, {
			Name:        "repo",
			Usage:       "create or configure repository",
			Subcommands: repo.Commands(),
		}},
	}

	if err := app.Run(os.Args); err != nil {
		log.New().Fatal(err)
	}
}
