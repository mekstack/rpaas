package main

import (
	"fmt"
	"github.com/mekstack/nataas/core/internal/config"
	"github.com/mekstack/nataas/core/internal/controller"
	"github.com/mekstack/nataas/core/internal/grpc_api/domain_service"
	"github.com/mekstack/nataas/core/internal/grpc_api/project_service"
	"github.com/mekstack/nataas/core/internal/grpc_api/subdomain_service"
	"github.com/mekstack/nataas/core/internal/storage"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	appConfig := config.MustConfig()

	store := storage.MustConnect(
		appConfig.Redis.Host,
		appConfig.Redis.Port,
		appConfig.Redis.UserName,
		appConfig.Redis.Password,
	)

	cnt := controller.New(store)

	listener, err := net.Listen(
		"tcp",
		fmt.Sprintf("%s:%d", appConfig.GrpcServer.Host, appConfig.GrpcServer.Port),
	)

	if err != nil {
		log.Fatal(err.Error())
	}

	grpcServer := grpc.NewServer()

	domain_service.Register(grpcServer, cnt)
	subdomain_service.Register(grpcServer, cnt)
	project_service.Register(grpcServer, cnt)

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatal(err.Error())
	}

}
