package start

import (
	"errors"

	"github.com/urfave/cli/v2"

	logutil "github.com/lthibault/ipfs-fileshare/internal/util/log"
	fserve "github.com/lthibault/ipfs-fileshare/pkg"
)

// Flags for the `start` command
func Flags() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:    "repo",
			Aliases: []string{"r"},
			Usage:   "path to IPFS repository",
			EnvVars: []string{"WW_REPO"},
		},
	}
}

// Run the `start` command
func Run() cli.ActionFunc {
	return func(c *cli.Context) error {
		s := fserve.New(
			fserve.WithLogger(logutil.New(c)),
		)

		if err := s.Start(); err != nil {
			return err
		}
		defer s.Close()

		return errors.New("TODO:  do something with server")
	}
}
