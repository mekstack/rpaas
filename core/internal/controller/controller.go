package controller

import (
	"context"
	"fmt"
	"github.com/mekstack/nataas/core/internal/storage"
	proto "github.com/mekstack/nataas/core/proto/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net"
	"strconv"
	"strings"
)

type Controller interface {
	GetDomainsPool(context.Context) ([]*proto.Domain, error)
	GetOccupiedSubdomains(context.Context) ([]*proto.Subdomain, error)
	GetProjectInfo(context.Context, uint32) (*proto.Project, error)
	AddProjectInfo(context.Context, uint32, *proto.Route) (*proto.Project, error)
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

	projectRoutes, err := (*store).GetProjectRoutes(ctx, projectNumber)
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

func (c *controller) AddProjectInfo(ctx context.Context, projectCode uint32, route *proto.Route) (*proto.Project, error) {
	subdomains := make([]string, 0)

	if nil == net.ParseIP(route.TargetIp) {
		return nil, status.Errorf(
			codes.InvalidArgument,
			"Invalid Host",
		)
	}

	if route.Port > 65_535 || route.Port <= 0 {
		return nil, status.Errorf(
			codes.InvalidArgument,
			"Invalid Port",
		)
	}

	for _, subdomain := range route.Subdomains {
		subdomain := subdomain.Name
		isValid, err := c.isValidSubdomain(ctx, subdomain)

		if err != nil {
			return nil, err
		}

		if !isValid {
			return nil, status.Errorf(
				codes.InvalidArgument,
				fmt.Sprintf("Subdomain %s is not valid", subdomain),
			)
		}

		subdomains = append(subdomains, subdomain)
	}

	err := (*c.store).AddRouteToProject(ctx, projectCode, route.TargetIp, route.Port, subdomains)
	if err != nil {
		return nil, err
	}

	return c.GetProjectInfo(ctx, projectCode)
}

func (c *controller) isValidSubdomain(ctx context.Context, subdomain string) (bool, error) {
	validDomains, err := (*c.store).GetDomainsPool(ctx)
	if err != nil {
		return false, err
	}

	partsOfSubdomain := strings.Split(subdomain, ".")
	domainOfSubdomain := strings.Join(
		partsOfSubdomain[len(partsOfSubdomain)-2:],
		".",
	)
	for _, validDomain := range validDomains {

		if validDomain == domainOfSubdomain {
			return true, nil
		}
	}

	return false, err
}

func New(store storage.Storage) *controller {
	return &controller{store: &store}
}
