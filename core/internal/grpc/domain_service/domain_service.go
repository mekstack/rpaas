package domain_service

import (
	"context"

	proto "github.com/mekstack/nataas/core/proto/pb"

	"google.golang.org/grpc"
)

type DomainController interface {
	GetDomainsPool(context.Context) ([]*proto.Domain, error)
}

type service struct {
	proto.UnimplementedDomainServiceServer
	domainController DomainController
}

func Register(server *grpc.Server, controller DomainController) {
	proto.RegisterDomainServiceServer(server, &service{
		domainController: controller,
	})
}

func (d *service) GetDomainsPool(ctx context.Context, request *proto.GetDomainsPoolRequest) (*proto.GetDomainsPoolResponse, error) {
	pool, err := d.domainController.GetDomainsPool(ctx)

	if err != nil {
		return nil, err
	}

	return &proto.GetDomainsPoolResponse{
		Domains: pool,
	}, nil
}
