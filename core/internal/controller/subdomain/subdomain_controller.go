package subdomain_controller

import (
	"context"

	proto "github.com/mekstack/nataas/core/proto/pb"

	"go.uber.org/zap"
)

type SubdomainProvider interface {
	GetOccupiedSubdomains(context.Context) ([]string, error)
	AddToOccupiedSubdomains(context.Context, string) error
}

type controller struct {
	log               *zap.Logger
	subdomainProvider SubdomainProvider
}

func New(provider SubdomainProvider, logger *zap.Logger) *controller {
	return &controller{
		log:               logger,
		subdomainProvider: provider,
	}
}

func (c *controller) GetOccupiedSubdomains(ctx context.Context) ([]*proto.Subdomain, error) {
	pool, err := c.subdomainProvider.GetOccupiedSubdomains(ctx)
	if err != nil {
		return nil, err
	}

	subdomains := make([]*proto.Subdomain, len(pool))
	for i, subdomainName := range pool {
		subdomains[i] = &proto.Subdomain{
			Name: subdomainName,
		}
	}

	return subdomains, nil
}
