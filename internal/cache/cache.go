package xdscache

import (
	"context"
	"errors"
	"golang.org/x/sync/errgroup"
	"log/slog"
	"sync"
	certclient "xds_server/internal/client/cert"
	routesclient "xds_server/internal/client/routes"
	xdsconverter "xds_server/internal/converter"
	"xds_server/internal/snapshot"

	"github.com/envoyproxy/go-control-plane/pkg/cache/types"
	"github.com/envoyproxy/go-control-plane/pkg/cache/v3"
	"github.com/envoyproxy/go-control-plane/pkg/resource/v3"
)

func InitializeCache(ctx context.Context, certClient certclient.Client, routesClient routesclient.Client, logger *slog.Logger) (cdsCache *cache.LinearCache, ldsCache *cache.LinearCache, rdsCache *cache.LinearCache, err error) {
	mc, ml, mr := make(map[string]types.Resource), make(map[string]types.Resource), make(map[string]types.Resource)
	var clusterMutex, listenerMutex, routeMutex sync.Mutex

	domainEpCh, errCh := routesClient.Routes(ctx)
	select {
	case err = <-errCh:
		if err != nil {
			return nil, nil, nil, err
		}
	default:
	}

	g, ctx := errgroup.WithContext(ctx)
	for domainEp := range domainEpCh {
		//LISTENER WITH TLS
		g.Go(func() error {
			cert, err := certClient.Cert(ctx, domainEp.Domain)
			if err != nil {
				if errors.Is(err, certclient.ErrorNoData) {
					return nil
				}
				logger.Error("failed to set certificate", slog.String("domain", domainEp.Domain), slog.String("error", err.Error()))
				return err
			}

			listenerSafe := xdsconverter.ListenerConverter(domainEp, cert)

			_, key, value := snapshot.GenerateListener(listenerSafe)

			listenerMutex.Lock()
			ml[key] = value
			listenerMutex.Unlock()
			return nil
		})

		//UNSAFE LISTENER
		g.Go(func() error {
			listenerUnsafe := xdsconverter.ListenerConverter(domainEp, nil)

			_, key, value := snapshot.GenerateListener(listenerUnsafe)

			listenerMutex.Lock()
			ml[key] = value
			listenerMutex.Unlock()

			return nil
		})

		//VIRTUAL HOSTS
		g.Go(func() error {
			vh := xdsconverter.VirtualHostConverter(domainEp)
			_, key, value := snapshot.GenerateRoute(vh)

			routeMutex.Lock()
			mr[key] = value
			routeMutex.Unlock()

			return nil
		})

		//CLUSTER
		g.Go(func() error {
			clusters := xdsconverter.ClusterConverter(domainEp)

			_, key, value := snapshot.GenerateCluster(clusters)

			clusterMutex.Lock()
			mc[key] = value
			clusterMutex.Unlock()

			return nil
		})

	}

	if err := g.Wait(); err != nil {
		slog.Error("failed to register domains", slog.String("error", err.Error()))
		return nil, nil, nil, err
	}

	cdsOptions := cache.WithInitialResources(mc)
	ldsOptions := cache.WithInitialResources(ml)
	rdsOptions := cache.WithInitialResources(mr)

	cdsCache = cache.NewLinearCache(resource.ClusterType, cdsOptions)
	ldsCache = cache.NewLinearCache(resource.ListenerType, ldsOptions)
	rdsCache = cache.NewLinearCache(resource.RouteType, rdsOptions)

	return
}
