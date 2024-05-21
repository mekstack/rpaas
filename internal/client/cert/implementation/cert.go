package certclientimpl

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log/slog"
	certclient "xds_server/internal/client/cert"
	xdsmodels "xds_server/internal/models"
	"xds_server/pkg/cert"
)

func (c *client) Cert(ctx context.Context, domain string) (*xdsmodels.Cert, error) {
	res, err := c.api.Cert(ctx, &cert.CertRequest{
		Domain: domain,
	})
	if err != nil {
		if status.Convert(err).Code() == codes.NotFound {
			c.logger.Info("no certificate data", slog.String("domain", domain))
			return nil, certclient.ErrorNoData
		}
		c.logger.Error("failed to get cert", slog.String("domain", domain), slog.String("error", err.Error()))
		return nil, err

	}

	return &xdsmodels.Cert{
		CertificateChain: string(res.CertificateChain),
		PrivateKey:       string(res.PrivateKey),
	}, nil
}
