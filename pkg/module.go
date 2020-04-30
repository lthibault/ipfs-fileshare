package fshare

import (
	"context"

	"go.uber.org/fx"

	"github.com/ipfs/go-ipfs/core"
	"github.com/ipfs/go-ipfs/core/coreapi"
	"github.com/ipfs/go-ipfs/core/node/libp2p"
	"github.com/ipfs/go-ipfs/repo"
	iface "github.com/ipfs/interface-go-ipfs-core"
	log "github.com/lthibault/log/pkg"
)

func server(svc *Server, opt []Option) fx.Option {
	return fx.Options(
		fx.NopLogger,
		fx.Supply(opt),
		fx.Provide(
			newCtx,
			userConfig,
			newRepository,
			newBuildCfg,
			core.NewNode,
			coreapi.NewCoreAPI,
			newServer,
		),
		fx.Populate(svc),
	)
}

func newCtx(lx fx.Lifecycle) context.Context {
	ctx, cancel := context.WithCancel(context.Background())
	lx.Append(fx.Hook{
		OnStop: func(context.Context) error {
			cancel()
			return nil
		},
	})

	return ctx
}

type userConfigOut struct {
	fx.Out

	Log      log.Logger
	RepoPath string `name:"path"`
}

func userConfig(opt []Option) (out userConfigOut, err error) {
	cfg := new(Config)
	for _, f := range withDefault(opt) {
		if err = f(cfg); err != nil {
			return
		}
	}

	out.Log = cfg.log
	out.RepoPath = cfg.repoPath
	return
}

func newBuildCfg(repo repo.Repo) (*core.BuildCfg, error) {
	return &core.BuildCfg{
		Online:    true,
		Permanent: true,
		Routing:   libp2p.DHTOption,
		ExtraOpts: map[string]bool{
			"pubsub": true,
			// "ipnsps": false,
			// "mplex":  false,
		},
		Repo: repo,
	}, nil
}

type serverConfig struct {
	fx.In

	CoreAPI iface.CoreAPI
	IPFS    *core.IpfsNode
}

func newServer(cfg serverConfig) Server {
	return Server{
		node: cfg.IPFS,
		api:  cfg.CoreAPI,
	}
}
