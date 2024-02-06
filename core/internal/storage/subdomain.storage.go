package storage

import (
	"context"
)

const subdomainTableKey = "subdomains"

func (s *storage) GetOccupiedSubdomains(ctx context.Context) ([]string, error) {
	occupiedSubdomains := make([]string, 0)

	request := s.db.SMembers(ctx, subdomainTableKey)
	if err := request.Err(); err != nil {
		return nil, err
	}

	for _, subDomainName := range request.Val() {
		occupiedSubdomains = append(occupiedSubdomains, subDomainName)
	}

	return occupiedSubdomains, nil
}

func (s *storage) AddToOccupiedSubdomains(ctx context.Context, subdomain string) error {
	request := s.db.SAdd(ctx, subdomainTableKey, subdomain)
	return request.Err()
}
