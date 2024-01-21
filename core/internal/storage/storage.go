package storage

import (
	"context"
	"fmt"
	proto "github.com/mekstack/nataas/core/proto/pb"
	"github.com/redis/go-redis/v9"
)

type Storage interface {
	GetDomainsPool(ctx context.Context) ([]*proto.Domain, error)
}

type storage struct {
	db *redis.Client
}

func (s *storage) GetDomainsPool(ctx context.Context) ([]*proto.Domain, error) {
	availableDomains := make([]*proto.Domain, 0)
	for _, domainName := range s.db.SMembers(ctx, "domain").Val() {
		availableDomains = append(availableDomains, &proto.Domain{
			Name: domainName,
		})
	}
	return availableDomains, nil
}

func New(host string, port uint) *storage {
	return &storage{
		db: redis.NewClient(&redis.Options{
			Addr:     fmt.Sprintf("%s:%d", host, port),
			Password: "",
			DB:       0,
		}),
	}
}
