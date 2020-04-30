// Package logutil contains shared utilities for configuring loggers from a cli context.
package logutil

import (
	log "github.com/lthibault/log/pkg"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

// New logger from a cli context
func New(c *cli.Context) log.Logger {
	return log.New(WithLevel(c), WithFormat(c))
}

// WithLevel returns a log.Option that configures a logger's level.
func WithLevel(c *cli.Context) (opt log.Option) {
	var level = log.NullLevel
	defer func() {
		opt = log.OptLevel(level)
	}()

	if c.String("logfmt") == "none" {
		return
	}

	switch c.String("loglvl") {
	case "trace", "t":
		level = log.TraceLevel
	case "debug", "d":
		level = log.DebugLevel
	case "info", "i":
		level = log.InfoLevel
	case "warn", "warning", "w":
		level = log.WarnLevel
	case "error", "err", "e":
		level = log.ErrorLevel
	case "fatal", "f":
		level = log.FatalLevel
	default:
		level = log.InfoLevel
	}

	return
}

// WithFormat returns an option that configures a logger's format.
func WithFormat(c *cli.Context) log.Option {
	var fmt logrus.Formatter

	switch c.String("logfmt") {
	case "none":
	case "json":
		fmt = &logrus.JSONFormatter{PrettyPrint: c.Bool("prettyprint")}
	default:
		fmt = new(logrus.TextFormatter)
	}

	return log.OptFormatter(fmt)
}
