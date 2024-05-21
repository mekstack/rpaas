package app

import (
	"context"
	"fmt"
	"log/slog"
	"net"

	clusterservice "github.com/envoyproxy/go-control-plane/envoy/service/cluster/v3"
	listenerservice "github.com/envoyproxy/go-control-plane/envoy/service/listener/v3"
	routeservice "github.com/envoyproxy/go-control-plane/envoy/service/route/v3"
	"github.com/envoyproxy/go-control-plane/pkg/server/v3"
	"github.com/envoyproxy/go-control-plane/pkg/test/v3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

func (a *App) initXDSImpl(ctx context.Context) error {
	cb := &test.Callbacks{Debug: true}
	cdsCache, ldsCache, rdsCache, err := a.ServiceProvider.Cache(ctx)
	if err != nil {
		return err
	}

	a.cdsImpl = server.NewServer(ctx, cdsCache, cb)
	a.ldsImpl = server.NewServer(ctx, ldsCache, cb)
	a.rdsImpl = server.NewServer(ctx, rdsCache, cb)

	return nil
}

func (a *App) initXDSGRPCServer(_ context.Context) error {
	var grpcOptions []grpc.ServerOption
	grpcOptions = append(grpcOptions,
		grpc.MaxConcurrentStreams(grpcMaxConcurrentStreams),
		grpc.KeepaliveParams(keepalive.ServerParameters{
			Time:    grpcKeepaliveTime,
			Timeout: grpcKeepaliveTimeout,
		}),
		grpc.KeepaliveEnforcementPolicy(keepalive.EnforcementPolicy{
			MinTime:             grpcKeepaliveMinTime,
			PermitWithoutStream: true,
		}),
	)
	a.xdsGRPC = grpc.NewServer(grpcOptions...)

	clusterservice.RegisterClusterDiscoveryServiceServer(a.xdsGRPC, a.cdsImpl)
	listenerservice.RegisterListenerDiscoveryServiceServer(a.xdsGRPC, a.ldsImpl)
	routeservice.RegisterRouteDiscoveryServiceServer(a.xdsGRPC, a.rdsImpl)

	return nil
}

func (a *App) initXDSDeps(ctx context.Context) error {
	deps := []func(ctx context.Context) error{
		a.initXDSImpl,
		a.initXDSGRPCServer,
	}

	for _, f := range deps {
		if err := f(ctx); err != nil {
			return err
		}
	}

	return nil
}

func (a *App) runXDSGRPCServer() error {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", a.cfg.XdsPort))
	if err != nil {
		a.logger.Error("failed to create connection for xds app", slog.String("error", err.Error()))
		return err
	}

	a.logger.Info("xds app started", slog.Uint64("port", uint64(a.cfg.XdsPort)))
	if err = a.xdsGRPC.Serve(listener); err != nil {
		a.logger.Error("failed to serve xds app", slog.String("error", err.Error()))
		return err
	}

	return nil
}
