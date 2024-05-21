package xdsservice

import (
	"context"
	"xds_server/internal/models"
)

type Service interface {
	ApplyEndpoint(context.Context, *xdsmodels.DomainEndpoint) error
}
