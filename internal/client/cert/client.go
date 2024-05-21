package certclient

import (
	"context"
	xdsmodels "xds_server/internal/models"
)

type Client interface {
	Cert(ctx context.Context, domain string) (*xdsmodels.Cert, error)
	//CertStream(ctx context.Context, domains chan xdsmodels.DomainEndpoint) (chan xdsmodels.Cert, error)
	CloseConn() error
}
