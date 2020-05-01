package fshare

import (
	"context"

	log "github.com/lthibault/log/pkg"
	"go.uber.org/fx"

	"github.com/ipfs/go-ipfs/core"
	iface "github.com/ipfs/interface-go-ipfs-core"
)

// Server exposes local files over IPFS, and responds to search queries.
type Server struct {
	app *fx.App

	log  log.Logger
	node *core.IpfsNode
	api  iface.CoreAPI
}

// New file-sharing service.
func New(opt ...Option) Server {
	var svc Server
	app := fx.New(server(&svc, opt))

	svc.app = app

	return svc
}

// Start serving files
func (svc Server) Start() error {
	return svc.app.Start(context.Background())
}

// Close the service
func (svc Server) Close() error {
	return svc.app.Stop(context.Background())
}

// Share a path over IPFS recursively.  Hidden files are ignored.
func (svc Server) Share(path string) error {
	return nil
	// return errors.New("Server.Share NOT IMPLEMENTED")
}
