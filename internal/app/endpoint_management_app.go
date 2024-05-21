package app

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log/slog"
	"net"
	"xds_server/internal/api"
	pb "xds_server/pkg/endpoint_management"
)

// initialization endpoint management server
func (a *App) initEMGRPCServer(_ context.Context) error {
	a.emGRPC = grpc.NewServer()

	reflection.Register(a.emGRPC)

	pb.RegisterEndpointManagementServer(a.emGRPC, a.emImpl)
	return nil
}

func (a *App) initEMImplementation(ctx context.Context) error {
	srv, err := a.ServiceProvider.Service(ctx)
	if err != nil {
		return err
	}

	a.emImpl = api.New(srv)
	return nil
}

func (a *App) initEMDeps(ctx context.Context) error {
	deps := []func(ctx context.Context) error{
		a.initEMImplementation,
		a.initEMGRPCServer,
	}

	for _, f := range deps {
		if err := f(ctx); err != nil {
			return err
		}
	}

	return nil
}

func (a *App) runEMGRPCServer() error {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", a.cfg.EndpointManagementPort))
	if err != nil {
		a.logger.Error("failed to create connection for endpoint_management app", slog.String("error", err.Error()))
		return err
	}
	a.logger.Info("endpoint_management app started", slog.Uint64("port", uint64(a.cfg.EndpointManagementPort)))

	if err = a.emGRPC.Serve(listener); err != nil {
		a.logger.Error("failed to serve endpoint_management app", slog.String("error", err.Error()))
		return err
	}
	return nil
}
