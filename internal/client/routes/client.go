package allroutesclient

import (
	"context"
	"xds_server/internal/models"
)

type Client interface {
	Routes(context.Context) (chan *xdsmodels.DomainEndpoint, chan error)
	CloseConn() error
}
