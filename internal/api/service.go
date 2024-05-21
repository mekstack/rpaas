package api

import (
	xdsservice "xds_server/internal/service"
	pb "xds_server/pkg/endpoint_management"
)

type Implementation struct {
	pb.UnimplementedEndpointManagementServer
	service xdsservice.Service
}

func New(service xdsservice.Service) *Implementation {
	return &Implementation{
		service: service,
	}
}
