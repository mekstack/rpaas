package storage

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"log"
)

type Storage interface {
	GetDomainsPool(context.Context) ([]string, error)
	GetOccupiedSubdomains(context.Context) ([]string, error)
	GetProjectRoutes(context.Context, uint32) ([]string, error)
	AddRouteToProject(context.Context, uint32, string, uint32, []string) error
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
	tableKey := "subdomains"
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

func (s *storage) GetProjectRoutes(ctx context.Context, projectCode uint32) ([]string, error) {
	tableKey := "project"

	request := s.db.SMembers(ctx, fmt.Sprintf("%s:%d", tableKey, projectCode))
	if err := request.Err(); err != nil {
		return nil, err
	}

	projectRoutes := make([]string, 0)
	for _, subdomainName := range request.Val() {
		projectRoutes = append(projectRoutes, subdomainName)
	}

	return projectRoutes, nil
}

func (s *storage) AddRouteToProject(
	ctx context.Context,
	projectCode uint32,
	targetIp string,
	port uint32,
	subdomains []string,
) error {
	tableKey := "project"
	projectRoutes := make([]string, 0)

	for _, subdomain := range subdomains {
		projectRoutes = append(projectRoutes, fmt.Sprintf("%s:%s:%d", subdomain, targetIp, port))
	}

	request := s.db.SAdd(ctx, fmt.Sprintf("%s:%d", tableKey, projectCode), projectRoutes)
	return request.Err()
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
