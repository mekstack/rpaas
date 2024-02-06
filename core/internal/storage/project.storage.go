package storage

import (
	"context"
	"fmt"
)

const projectTableKey = "project"

func (s *storage) GetProjectRoutes(ctx context.Context, projectCode uint32) ([]string, error) {

	request := s.db.SMembers(ctx, fmt.Sprintf("%s:%d", projectTableKey, projectCode))
	if err := request.Err(); err != nil {
		return nil, err
	}

	projectRoutes := make([]string, 0)
	for _, route := range request.Val() {
		projectRoutes = append(projectRoutes, route)
	}

	return projectRoutes, nil
}

func (s *storage) AddRouteToProject(
	ctx context.Context,
	projectCode uint32,
	endpoint, subdomainName string,
) error {
	request := s.db.SAdd(ctx, fmt.Sprintf("%s:%d", projectTableKey, projectCode), fmt.Sprintf("%s:%s", endpoint, subdomainName))
	if request.Err() != nil {
		return request.Err()
	}
	return s.AddSubdomain(ctx, subdomainName)
}
