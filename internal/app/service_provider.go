package app

import (
	"context"
	"log/slog"
	xdscache "xds_server/internal/cache"
	certclient "xds_server/internal/client/cert"
	certclientimpl "xds_server/internal/client/cert/implementation"
	routesclient "xds_server/internal/client/routes"
	routesclientimpl "xds_server/internal/client/routes/implementation"
	xdsservice "xds_server/internal/service"
	serviceimpl "xds_server/internal/service/implementation"

	"github.com/envoyproxy/go-control-plane/pkg/cache/v3"
)

type serviceProvider struct {
	cdsCache *cache.LinearCache
	ldsCache *cache.LinearCache
	rdsCache *cache.LinearCache

	logger *slog.Logger

	srv xdsservice.Service

	certCl       certclient.Client
	routesClient routesclient.Client

	nodeID          string
	certServerPort  int
	routeServerPort int
}

func (sp *serviceProvider) Service(ctx context.Context) (xdsservice.Service, error) {
	if sp.srv == nil {
		client, err := sp.CertClient()
		if err != nil {
			return nil, err
		}

		cdsCache, ldsCache, rdsCache, err := sp.Cache(ctx)
		if err != nil {
			return nil, err
		}

		sp.srv = serviceimpl.New(cdsCache, ldsCache, rdsCache, client, sp.logger, sp.nodeID)
	}

	return sp.srv, nil
}

func (sp *serviceProvider) CertClient() (certclient.Client, error) {
	if sp.certCl == nil {
		var err error
		sp.certCl, err = certclientimpl.New(sp.certServerPort, sp.logger)
		if err != nil {
			return nil, err
		}
	}
	return sp.certCl, nil
}

func (sp *serviceProvider) RoutesClient() (routesclient.Client, error) {
	if sp.routesClient == nil {
		var err error
		sp.routesClient, err = routesclientimpl.New(sp.routeServerPort, sp.logger)
		if err != nil {
			return nil, err
		}
	}
	return sp.routesClient, nil
}

func (sp *serviceProvider) Cache(ctx context.Context) (*cache.LinearCache, *cache.LinearCache, *cache.LinearCache, error) {
	if sp.cdsCache == nil || sp.ldsCache == nil || sp.rdsCache == nil {
		certClient, err := sp.CertClient()
		if err != nil {
			return nil, nil, nil, err
		}

		routesClient, err := sp.RoutesClient()
		if err != nil {
			return nil, nil, nil, err
		}

		sp.cdsCache, sp.ldsCache, sp.rdsCache, err = xdscache.InitializeCache(ctx, certClient, routesClient, sp.logger)
		if err != nil {
			return nil, nil, nil, err
		}
	}
	return sp.cdsCache, sp.ldsCache, sp.rdsCache, nil
}

func newServiceProvider(logger *slog.Logger, nodeID string, certServerPort, routeServerPort int) *serviceProvider {
	return &serviceProvider{
		logger:          logger,
		nodeID:          nodeID,
		certServerPort:  certServerPort,
		routeServerPort: routeServerPort,
	}
}
