package main

import (
	"fmt"
	"github.com/mekstack/nataas/core/internal/config"
	"github.com/mekstack/nataas/core/internal/controller"
	"github.com/mekstack/nataas/core/internal/grpc_api/domain_service"
	"github.com/mekstack/nataas/core/internal/storage"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	appConfig := config.MustRun()

	listener, err := net.Listen(
		"tcp",
		fmt.Sprintf("%s:%d", appConfig.GrpcServerHost, appConfig.GrpcServerPort),
	)

	if err != nil {
		log.Fatal("Something went wrong", err.Error())
	}

	store := storage.New(appConfig.RedisHost, appConfig.RedisPort)

	cnt := controller.New(store)

	grpcServer := grpc.NewServer()

	domain_service.Register(grpcServer, cnt)

	if err := grpcServer.Serve(listener); err != nil {
		return
	}

}
