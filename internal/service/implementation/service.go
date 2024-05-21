package serviceimpl

import (
	"github.com/go-playground/validator/v10"
	"log/slog"
	"sync"
	certclient "xds_server/internal/client/cert"
	xdsservice "xds_server/internal/service"

	"github.com/envoyproxy/go-control-plane/pkg/cache/v3"
)

type service struct {
	cdsCache *cache.LinearCache
	ldsCache *cache.LinearCache
	rdsCache *cache.LinearCache

	mu sync.Mutex

	v *validator.Validate

	certClient certclient.Client

	logger *slog.Logger

	nodeID string
}

func New(cdsCache, ldsCache, rdsCache *cache.LinearCache, certClient certclient.Client, logger *slog.Logger, nodeID string) xdsservice.Service {
	return &service{
		cdsCache:   cdsCache,
		ldsCache:   ldsCache,
		rdsCache:   rdsCache,
		v:          validator.New(),
		certClient: certClient,
		logger:     logger,
		nodeID:     nodeID,
	}
}
