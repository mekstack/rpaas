package storage

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

type storage struct {
	db  *redis.Client
	log *zap.Logger
}

func MustConnect(addr string, logger *zap.Logger) *storage {
	storage := &storage{
		db: redis.NewClient(&redis.Options{
			Addr: addr,
		}),
		log: logger,
	}

	err := storage.db.Ping(context.Background()).Err()
	if err != nil {
		panic(fmt.Errorf("Redis connection error: %s", err))
	}

	return storage
}
