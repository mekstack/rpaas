package allroutesclientimpl

import (
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
	"io"
	"log/slog"
	"sync"
	"xds_server/internal/models"
)

const (
	domainEndpointChBuffer = 5
)

func (c *client) Routes(ctx context.Context) (chan *xdsmodels.DomainEndpoint, chan error) {
	domainEps := make(chan *xdsmodels.DomainEndpoint, domainEndpointChBuffer)
	errCh := make(chan error, 1)

	routes, err := c.api.AllRoutes(ctx, &emptypb.Empty{})
	if err != nil {
		c.logger.Error("failed to get routes", slog.String("error", err.Error()))

		errCh <- err
		close(domainEps)
		close(errCh)

		return nil, errCh
	}

	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(domainEps)
		defer close(errCh)

		for {
			res, err := routes.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				errCh <- err
				return
			}

			domainEps <- &xdsmodels.DomainEndpoint{
				Domain: res.GetDomain(),
				Host:   res.GetHost(),
				Port:   res.GetPort(),
			}
		}
	}()

	return domainEps, errCh
}
