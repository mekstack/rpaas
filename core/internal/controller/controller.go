package controller

import (
	"context"
	"github.com/mekstack/nataas/core/internal/storage"
	proto "github.com/mekstack/nataas/core/proto/pb"
	"strconv"
	"strings"
)

type Controller interface {
	GetDomainsPool(context.Context) ([]*proto.Domain, error)
	GetOccupiedSubdomains(context.Context) ([]*proto.Subdomain, error)
	GetProjectInfo(context.Context, uint32) (*proto.Project, error)
}

type controller struct {
	store *storage.Storage
}

func (c *controller) GetDomainsPool(ctx context.Context) (domains []*proto.Domain, err error) {
	store := c.store
	availableDomains, err := (*store).GetDomainsPool(ctx)
	if err != nil {
		return nil, err
	}

	for _, domain := range availableDomains {
		domains = append(domains, &proto.Domain{Name: domain})
	}

	return domains, nil
}

func (c *controller) GetOccupiedSubdomains(ctx context.Context) (subdomains []*proto.Subdomain, err error) {
	store := c.store
	occupiedSubdomains, err := (*store).GetOccupiedSubdomains(ctx)
	if err != nil {
		return nil, err
	}

	for _, subdomain := range occupiedSubdomains {
		subdomains = append(subdomains, &proto.Subdomain{Name: subdomain})
	}

	return subdomains, nil
}

func (c *controller) GetProjectInfo(ctx context.Context, projectNumber uint32) (*proto.Project, error) {
	store := c.store

	projectRoutes, err := (*store).GetProjectInfo(ctx, projectNumber)
	if err != nil {
		return nil, err
	}

	subdomains := make(map[string][]*proto.Subdomain)

	for _, route := range projectRoutes {
		route := strings.Split(route, ":")

		hostAddr := strings.Join(route[1:3], ":")
		_, ok := subdomains[hostAddr]
		if ok == false {
			subdomains[hostAddr] = make([]*proto.Subdomain, 0)
		}
		subdomains[hostAddr] = append(subdomains[hostAddr], &proto.Subdomain{
			Name: route[0],
		})

	}

	var routes = make([]*proto.Route, 0)

	for hostAddr, subdomains := range subdomains {
		hostAddr := strings.Split(hostAddr, ":")

		targetIp := hostAddr[0]
		port, err := strconv.Atoi(hostAddr[1])
		if err != nil {
			return nil, err
		}
		routes = append(routes, &proto.Route{
			TargetIp:   targetIp,
			Port:       uint32(port),
			Subdomains: subdomains,
		})
	}

	return &proto.Project{
		Code:   projectNumber,
		Routes: routes,
	}, nil
}

func New(store storage.Storage) *controller {
	return &controller{store: &store}
}
