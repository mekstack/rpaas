package controller

import (
	"context"
	"github.com/mekstack/nataas/core/internal/storage"
	proto "github.com/mekstack/nataas/core/proto/pb"
)

type Controller interface {
	GetDomainsPool(ctx context.Context) ([]*proto.Domain, error)
}

type controller struct {
	store *storage.Storage
}

func (c *controller) GetDomainsPool(ctx context.Context) ([]*proto.Domain, error) {
	store := c.store
	return (*store).GetDomainsPool(ctx)
}

func New(store storage.Storage) *controller {
	return &controller{store: &store}
}
