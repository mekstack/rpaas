package domain_service

import (
	"context"
	proto "github.com/mekstack/nataas/core/proto/pb"
	"google.golang.org/grpc"
)

type ServiceApi struct {
	proto.UnimplementedDomainServiceServer
}

func (s ServiceApi) GetDomainsPool(ctx context.Context, request *proto.GetDomainsPoolRequest) (*proto.GetDomainsPoolResponse, error) {
	return &proto.GetDomainsPoolResponse{
		Domains: []*proto.Domain{
			&proto.Domain{
				Name: "mekstack.ru",
			},
		},
	}, nil
}

func Register(server *grpc.Server) {
	proto.RegisterDomainServiceServer(server, ServiceApi{})
}
