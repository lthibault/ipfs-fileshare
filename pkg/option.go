package fshare

import (
	log "github.com/lthibault/log/pkg"
)

// Option type for Host
type Option func(*Config) error

// Config contains user-specified parameters.  These will be made available to Fx.
type Config struct {
	log      log.Logger
	repoPath string
}

// WithLogger sets the logger.
func WithLogger(l log.Logger) Option {
	if l == nil {
		l = log.New(log.OptLevel(log.FatalLevel))
	}

	return func(c *Config) (err error) {
		c.log = l
		return
	}
}

// WithRepoPath sets the repository path.
func WithRepoPath(path string) Option {
	return func(c *Config) (err error) {
		c.repoPath = path
		return
	}
}

func withDefault(opt []Option) []Option {
	return append([]Option{
		WithLogger(nil),
		WithRepoPath(""),
	}, opt...)
}
