package certclientimpl

import (
	"context"
	"errors"
	"golang.org/x/sync/errgroup"
	"io"
	xdsmodels "xds_server/internal/models"
	pb "xds_server/pkg/cert"
)

func (c *client) CertStream(ctx context.Context, domainEps chan xdsmodels.DomainEndpoint) ([]xdsmodels.Cert, error) {
	certificates := make([]xdsmodels.Cert, 0, len(domainEps))

	cl, err := c.api.CertStream(ctx)
	if err != nil {
		return nil, err
	}

	g, _ := errgroup.WithContext(ctx)

	for domainEp := range domainEps {
		g.Go(func() error {
			if err = cl.Send(&pb.CertRequest{
				Domain: domainEp.Domain,
			}); err != nil && !errors.Is(err, io.EOF) {
				return err
			}

			res, err := cl.Recv()
			if err != nil && err != io.EOF {
				return err
			}

			certificates = append(certificates, xdsmodels.Cert{
				CertificateChain: string(res.GetCertificateChain()),
				PrivateKey:       string(res.GetPrivateKey()),
			})

			return nil
		})
	}

	if err = g.Wait(); err != nil {
		return nil, err
	}

	if err = cl.CloseSend(); err != nil {
		return nil, err
	}

	return certificates, nil
}
