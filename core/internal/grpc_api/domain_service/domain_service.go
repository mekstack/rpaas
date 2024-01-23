package domain_service

import (
	"context"
	"github.com/mekstack/nataas/core/internal/controller"
	proto "github.com/mekstack/nataas/core/proto/pb"
	"google.golang.org/grpc"
)

type DomainApi struct {
	cnt *controller.Controller
	proto.UnimplementedDomainServiceServer
}

func (d *DomainApi) GetDomainsPool(ctx context.Context, request *proto.GetDomainsPoolRequest) (*proto.GetDomainsPoolResponse, error) {
	availableDomains, err := (*d.cnt).GetDomainsPool(ctx)

	if err != nil {
		return nil, err
	}

	return &proto.GetDomainsPoolResponse{
		Domains: availableDomains,
	}, nil
}

func Register(server *grpc.Server, controller controller.Controller) {
	proto.RegisterDomainServiceServer(server, &DomainApi{
		cnt: &controller,
	})
}
