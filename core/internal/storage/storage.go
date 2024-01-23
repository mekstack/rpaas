package storage

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"log"
)

type Storage interface {
	GetDomainsPool(ctx context.Context) ([]string, error)
	GetOccupiedSubdomains(ctx context.Context) ([]string, error)
	GetProjectInfo(context.Context, uint32) ([]string, error)
}

type storage struct {
	db *redis.Client
}

func (s *storage) GetDomainsPool(ctx context.Context) ([]string, error) {
	tableKey := "domains"
	availableDomains := make([]string, 0)

	request := s.db.SMembers(ctx, tableKey)
	if err := request.Err(); err != nil {
		return nil, err
	}

	for _, domainName := range request.Val() {
		availableDomains = append(availableDomains, domainName)
	}

	return availableDomains, nil
}

func (s *storage) GetOccupiedSubdomains(ctx context.Context) ([]string, error) {
	tableKey := "domains"
	occupiedSubdomains := make([]string, 0)

	request := s.db.SMembers(ctx, tableKey)
	if err := request.Err(); err != nil {
		return nil, err
	}

	for _, subDomainName := range request.Val() {
		occupiedSubdomains = append(occupiedSubdomains, subDomainName)
	}

	return occupiedSubdomains, nil
}

func (s *storage) GetProjectInfo(ctx context.Context, projectNumber uint32) ([]string, error) {
	tableKey := "project"
	projectRoutes := make([]string, 0)

	request := s.db.SMembers(ctx, fmt.Sprintf("%s:%d", tableKey, projectNumber))
	if err := request.Err(); err != nil {
		return nil, err
	}

	for _, subDomainName := range request.Val() {
		projectRoutes = append(projectRoutes, subDomainName)
	}

	return projectRoutes, nil
}

func MustConnect(host string, port uint, userName, password string) *storage {
	storage := &storage{
		db: redis.NewClient(&redis.Options{
			Addr:     fmt.Sprintf("%s:%d", host, port),
			Username: userName,
			Password: password,
			DB:       0,
		}),
	}
	err := storage.db.Ping(context.Background()).Err()
	if err != nil {
		log.Fatal(err.Error())
	}
	return storage
}
