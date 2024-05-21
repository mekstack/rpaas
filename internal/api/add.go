package api

import (
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
	xdsmodels "xds_server/internal/models"
	pb "xds_server/pkg/endpoint_management"
)

func (i *Implementation) Add(ctx context.Context, req *pb.Endpoint) (*emptypb.Empty, error) {
	domainEp := &xdsmodels.DomainEndpoint{
		Domain: req.GetDomain(),
		Host:   req.GetHost(),
		Port:   req.GetPort(),
	}

	if err := i.service.ApplyEndpoint(ctx, domainEp); err != nil {
		return nil, err
	}

	return nil, nil
}
