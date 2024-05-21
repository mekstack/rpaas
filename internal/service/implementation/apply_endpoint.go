package serviceimpl

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"log/slog"
	xdsconverter "xds_server/internal/converter"
	"xds_server/internal/models"
	"xds_server/internal/snapshot"
)

func (s *service) ApplyEndpoint(ctx context.Context, domainEp *xdsmodels.DomainEndpoint) error {
	err := s.validateEndpoint(domainEp)
	if err != nil {
		return err
	}

	g, ctx := errgroup.WithContext(ctx)

	// CLUSTER
	g.Go(func() error {
		cl := xdsconverter.ClusterConverter(domainEp)
		resource, _, _ := snapshot.GenerateCluster(cl)

		s.mu.Lock()
		if s.cdsCache, err = s.updateCache(s.cdsCache, resource); err != nil {
			s.logger.Error("failed to update cluster cache", slog.String("error", err.Error()))
			return fmt.Errorf("cluster cache update error: %s", err.Error())
		}
		s.mu.Unlock()

		s.logger.Info("cluster update success", slog.String("Domain", domainEp.Domain))

		return nil
	})

	// ROUTE
	g.Go(func() error {
		vh := xdsconverter.VirtualHostConverter(domainEp)

		resource, _, _ := snapshot.GenerateRoute(vh)

		s.mu.Lock()
		if s.rdsCache, err = s.updateCache(s.rdsCache, resource); err != nil {
			s.logger.Error("failed to update route cache", slog.String("error", err.Error()))
			return fmt.Errorf("route cache update error: %s", err.Error())
		}
		s.mu.Unlock()

		s.logger.Info("route update success", slog.String("Domain", domainEp.Domain))

		return nil
	})

	// SAFE LISTENER
	g.Go(func() error {
		cert, err := s.certClient.Cert(ctx, domainEp.Domain)
		if err != nil {
			s.logger.Error("failed to get certificates", slog.String("Domain", domainEp.Domain), slog.String("error", err.Error()))
			return fmt.Errorf("get certificates error: %s", err.Error())
		}

		listenerSafe := xdsconverter.ListenerConverter(domainEp, cert)

		s.mu.Lock()
		resource, _, _ := snapshot.GenerateListener(listenerSafe)
		if s.ldsCache, err = s.updateCache(s.ldsCache, resource); err != nil {
			s.logger.Error("failed to update listener cache", slog.String("error", err.Error()))
			return fmt.Errorf("listener cache update error: %s", err.Error())
		}
		s.mu.Unlock()

		s.logger.Info("https listener update success", slog.String("Domain", domainEp.Domain))

		return nil
	})

	// UNSAFE LISTENER
	g.Go(func() error {
		listenerUnsafe := xdsconverter.ListenerConverter(domainEp, nil)

		resource, _, _ := snapshot.GenerateListener(listenerUnsafe)

		s.mu.Lock()
		if s.ldsCache, err = s.updateCache(s.ldsCache, resource); err != nil {
			s.logger.Error("failed to update listener cache", slog.String("error", err.Error()))
			return fmt.Errorf("listener cache update error: %s", err.Error())
		}
		s.mu.Unlock()

		s.logger.Info("http listener update success", slog.String("Domain", domainEp.Domain))

		return nil
	})

	//WAITING FOR APPLYING
	if err = g.Wait(); err != nil {
		s.logger.Error("failed to apply endpoint", slog.String("error", err.Error()))
		return fmt.Errorf("apply endpoint error: %s", err.Error())
	}

	s.logger.Info("data applied", "endpoint", domainEp)

	return nil
}
