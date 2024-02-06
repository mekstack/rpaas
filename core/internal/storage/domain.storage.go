package storage

import "context"

const domainTableKey = "domains"

func (s *storage) GetDomainsPool(ctx context.Context) ([]string, error) {
	availableDomains := make([]string, 0)

	request := s.db.SMembers(ctx, domainTableKey)
	if err := request.Err(); err != nil {
		return nil, err
	}

	for _, domainName := range request.Val() {
		availableDomains = append(availableDomains, domainName)
	}

	return availableDomains, nil
}
