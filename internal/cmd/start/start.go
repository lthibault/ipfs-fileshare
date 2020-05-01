package start

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/radovskyb/watcher"
	"github.com/urfave/cli/v2"

	logutil "github.com/lthibault/ipfs-fileshare/internal/util/log"
	fshare "github.com/lthibault/ipfs-fileshare/pkg"
	log "github.com/lthibault/log/pkg"
)

var exit <-chan os.Signal

func init() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)
	exit = c
}

// Flags for the `start` command
func Flags() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:    "repo",
			Aliases: []string{"r"},
			Usage:   "path to IPFS repository",
			EnvVars: []string{"FSHARE_REPO"},
		},
		&cli.PathFlag{
			Name:    "dir",
			Aliases: []string{"d"},
			Usage:   "path to shared directory",
			Value:   "/tmp/fshare", // TODO:  pick something better
			EnvVars: []string{"FSHARE_DIR"},
		},
	}
}

// Run the `start` command
func Run() cli.ActionFunc {
	return func(c *cli.Context) error {
		log := logutil.New(c)

		// Set up file-sharing server.
		s := fshare.New(
			fshare.WithLogger(log),
		)

		if err := s.Start(); err != nil {
			return err
		}
		defer s.Close()

		switch err := os.MkdirAll(c.Path("dir"), 0750); err {
		case nil, os.ErrExist:
		default:
			return err
		}

		if err := s.Share(c.Path("dir")); err != nil {
			return err
		}

		return watch(log, c.Path("dir"), s)
	}
}

func watch(l log.Logger, path string, s fshare.Server) error {
	w := watcher.New()
	w.IgnoreHiddenFiles(true)
	w.AddRecursive(path)
	if err := w.AddRecursive(path); err != nil {
		return err
	}

	go func() {
		defer w.Close()

		for {
			select {
			case event := <-w.Event:
				handle(l, s, event)
			case err := <-w.Error:
				l.WithError(err).Debug("fswatcher failure")
			case <-exit:
				return
			}
		}
	}()

	l.Debug("starting watcher")
	defer l.Debug("watcher stopped")

	return w.Start(time.Millisecond * 200)
}

func handle(l log.Logger, s fshare.Server, event watcher.Event) {
	l.WithFields(log.F{
		"op":       event.Op,
		"path":     event.Path,
		"old_path": event.OldPath,
		"dir":      event.IsDir(),
	}).Debug("got fs event")

	switch event.Op {
	case watcher.Create:
	case watcher.Remove:
	case watcher.Rename:
	case watcher.Write:
	}
}
