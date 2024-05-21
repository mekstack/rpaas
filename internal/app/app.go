package app

import (
	"context"
	"golang.org/x/sync/errgroup"
	"log/slog"
	"time"
	"xds_server/internal/api"
	xdsconfig "xds_server/internal/config"

	"github.com/envoyproxy/go-control-plane/pkg/server/v3"
	"google.golang.org/grpc"
)

const (
	grpcKeepaliveTime        = 30 * time.Second
	grpcKeepaliveTimeout     = 5 * time.Second
	grpcKeepaliveMinTime     = 30 * time.Second
	grpcMaxConcurrentStreams = 1000000
)

type App struct {
	xdsGRPC *grpc.Server
	emGRPC  *grpc.Server

	cdsImpl server.Server
	ldsImpl server.Server
	rdsImpl server.Server

	emImpl *api.Implementation

	ServiceProvider *serviceProvider

	logger *slog.Logger

	nodeID string

	cfg *xdsconfig.Config
}

func (a *App) initConfig(configPath string) error {
	cfg, err := xdsconfig.New(configPath)
	if err != nil {
		return err
	}

	a.cfg = cfg

	return nil
}

func (a *App) initServiceProvider(_ context.Context) error {
	a.ServiceProvider = newServiceProvider(
		a.logger,
		a.nodeID,
		a.cfg.CertServerPort,
		a.cfg.RouteServerPort,
	)

	return nil
}

func (a *App) initDeps(ctx context.Context, configPath string) error {
	//CONFIG INITIALIZATION
	if err := a.initConfig(configPath); err != nil {
		return err
	}

	deps := []func(context.Context) error{
		a.initServiceProvider,
		a.initXDSDeps,
		a.initEMDeps,
	}

	for _, f := range deps {
		err := f(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *App) Run(ctx context.Context) error {
	g, _ := errgroup.WithContext(ctx)
	g.Go(func() error {
		return a.runXDSGRPCServer()
	})

	g.Go(func() error {
		return a.runEMGRPCServer()
	})

	if err := g.Wait(); err != nil {
		return err
	}

	return nil
}

func (a *App) Stop() {
	a.xdsGRPC.GracefulStop()
	a.emGRPC.GracefulStop()

	if a.ServiceProvider.certCl != nil {
		if err := a.ServiceProvider.certCl.CloseConn(); err != nil {
			a.logger.Error(err.Error())
		}
	}
}

func New(ctx context.Context, configPath string, logger *slog.Logger, nodeID string) (*App, error) {
	logger.Debug("creating new xds app")
	a := &App{
		logger: logger,
		nodeID: nodeID,
	}

	logger.Debug("initializing all dependencies for xds app")
	if err := a.initDeps(ctx, configPath); err != nil {
		return nil, err
	}

	logger.Debug("all dependencies for xds app initialized")
	return a, nil
}
