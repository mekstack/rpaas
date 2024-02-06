package subdomain_service

import (
	"context"

	proto "github.com/mekstack/nataas/core/proto/pb"

	"google.golang.org/grpc"
)

type SubdomainController interface {
	GetOccupiedSubdomains(context.Context) ([]*proto.Subdomain, error)
}

type service struct {
	proto.UnimplementedSubdomainServiceServer
	subdomainController SubdomainController
}

func Register(server *grpc.Server, controller SubdomainController) {
	proto.RegisterSubdomainServiceServer(server, &service{
		subdomainController: controller,
	})
}

func (s *service) GetOccupiedSubdomains(ctx context.Context, request *proto.GetOccupiedSubdomainsRequest) (*proto.GetOccupiedSubdomainsResponse, error) {
	pool, err := s.subdomainController.GetOccupiedSubdomains(ctx)

	if err != nil {
		return nil, err
	}

	return &proto.GetOccupiedSubdomainsResponse{
		Subdomains: pool,
	}, nil
}
