package main

import (
	"fmt"
	"net"

	"github.com/mekstack/nataas/core/internal/config"
	domain_controller "github.com/mekstack/nataas/core/internal/controller/domain"
	project_controller "github.com/mekstack/nataas/core/internal/controller/project"
	subdomain_controller "github.com/mekstack/nataas/core/internal/controller/subdomain"
	"github.com/mekstack/nataas/core/internal/grpc/domain_service"
	"github.com/mekstack/nataas/core/internal/grpc/project_service"
	"github.com/mekstack/nataas/core/internal/grpc/subdomain_service"
	"github.com/mekstack/nataas/core/internal/storage"

	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func main() {
	conf := config.MustConfig()

	log := mustSetUpLogger(conf)
	defer log.Sync()

	log.Info("Config and logger was set up", zap.Any("Config", conf))

	store := storage.MustConnect(conf.Redis.Addr, log)

	lis, err := net.Listen("tcp", conf.GrpcServer.Addr)
	if err != nil {
		log.Sugar().Fatalln("Creation listener error:", err)
	}
	log.Info("Core start to listen:", zap.String("Addr", conf.GrpcServer.Addr))

	srv := grpc.NewServer()
	domain_service.Register(srv, domain_controller.New(store, log))
	subdomain_service.Register(srv, subdomain_controller.New(store, log))
	project_service.Register(srv, project_controller.New(store, log))

	if err := srv.Serve(lis); err != nil {
		log.Sugar().Fatalln("Serve error:", err)
	}

}

func mustSetUpLogger(conf *config.Config) *zap.Logger {
	var logger *zap.Logger
	var err error
	switch conf.Environment {
	case config.Development:
		logger, err = zap.NewDevelopment()
	case config.Production:
		logger, err = zap.NewProduction()
	default:
		panic("There is not valid Environment")
	}
	if err != nil {
		panic(fmt.Errorf("Logger not set: %s", err))
	}
	return logger
}
