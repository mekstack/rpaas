package domain_controller

import (
	"context"

	proto "github.com/mekstack/nataas/core/proto/pb"

	"go.uber.org/zap"
)

type DomainProvider interface {
	GetDomainsPool(context.Context) ([]string, error)
}

type controller struct {
	log            *zap.Logger
	domainProvider DomainProvider
}

func New(provider DomainProvider, logger *zap.Logger) *controller {
	return &controller{
		log:            logger,
		domainProvider: provider,
	}
}

func (c *controller) GetDomainsPool(ctx context.Context) ([]*proto.Domain, error) {
	pool, err := c.domainProvider.GetDomainsPool(ctx)
	if err != nil {
		return nil, err
	}

	domains := make([]*proto.Domain, len(pool))
	for i, domainName := range pool {
		domains[i] = &proto.Domain{
			Name: domainName,
		}
	}

	return domains, nil
}
