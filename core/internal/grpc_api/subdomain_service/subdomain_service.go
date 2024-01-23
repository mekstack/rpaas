package subdomain_service

import (
	"context"
	"github.com/mekstack/nataas/core/internal/controller"
	proto "github.com/mekstack/nataas/core/proto/pb"
	"google.golang.org/grpc"
)

type SubdomainApi struct {
	cnt *controller.Controller
	proto.UnimplementedSubdomainServiceServer
}

func (s *SubdomainApi) GetOccupiedSubdomains(ctx context.Context, request *proto.GetOccupiedSubdomainsRequest) (*proto.GetOccupiedSubdomainsResponse, error) {
	occupiedSubdomains, err := (*s.cnt).GetOccupiedSubdomains(ctx)

	if err != nil {
		return nil, err
	}

	return &proto.GetOccupiedSubdomainsResponse{
		Subdomains: occupiedSubdomains,
	}, nil
}

func Register(server *grpc.Server, controller controller.Controller) {
	proto.RegisterSubdomainServiceServer(server, &SubdomainApi{
		cnt: &controller,
	})
}
