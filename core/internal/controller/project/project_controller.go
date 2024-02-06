package project_controller

import (
	"context"
	"net"
	"strconv"
	"strings"

	domain_controller "github.com/mekstack/nataas/core/internal/controller/domain"
	subdomain_controller "github.com/mekstack/nataas/core/internal/controller/subdomain"
	proto "github.com/mekstack/nataas/core/proto/pb"

	"go.uber.org/zap"
)

type ProjectProvider interface {
	GetProjectRoutes(context.Context, uint32) ([]string, error)
	AddRouteToProject(context.Context, uint32, string, string) error
	subdomain_controller.SubdomainProvider
	domain_controller.DomainProvider
}

type controller struct {
	log             *zap.Logger
	projectProvider ProjectProvider
}

func New(provider ProjectProvider, logger *zap.Logger) *controller {
	return &controller{
		log:             logger,
		projectProvider: provider,
	}
}

func (c *controller) GetProject(ctx context.Context, projectCode uint32) (*proto.Project, error) {
	routes, err := c.projectProvider.GetProjectRoutes(ctx, projectCode)
	c.log.Debug("Project routes", zap.Uint32("ProjectCode", projectCode), zap.Any("Routes", routes))

	if err != nil {
		return nil, err
	}

	project := new(proto.Project)
	project.Code = projectCode
	project.Routes = make([]*proto.Route, len(routes))

	for i, route := range routes {
		sRoute := strings.Split(route, ":")
		if len(sRoute) != 3 {
			c.log.Warn("Incorrect route", zap.Error(ErrRouteNotValid))
			return nil, ErrRouteNotValid
		}
		endpoint := strings.Join(sRoute[:2], ":")
		subdomain := sRoute[2]

		project.Routes[i] = &proto.Route{
			Endpoint: endpoint,
			Subdomain: &proto.Subdomain{
				Name: subdomain,
			},
		}
	}

	return project, nil
}

func (c *controller) AddRouteToProject(ctx context.Context, projectCode uint32, route *proto.Route) (*proto.Project, error) {
	subdomain := route.Subdomain.Name
	if err := c.subdomainValidation(ctx, subdomain); err != nil {
		c.log.Warn("Subdomain validator", zap.Error(err))
		return nil, err
	}

	endpoint := route.Endpoint
	if err := c.endpointValidation(endpoint); err != nil {
		c.log.Warn("Endpoint validator", zap.Error(err))
		return nil, err
	}

	if err := c.projectProvider.AddRouteToProject(ctx, projectCode, endpoint, subdomain); err != nil {
		c.log.Warn("Add route to project", zap.Error(err))
		return nil, err
	}

	if err := c.projectProvider.AddToOccupiedSubdomains(ctx, subdomain); err != nil {
		c.log.Warn("Add subdomain to occupied subdomains", zap.Error(err))
		return nil, err
	}

	return c.GetProject(ctx, projectCode)
}

func (c *controller) subdomainValidation(ctx context.Context, subdomain string) error {
	// TODO: Add regex validator that subdomain
	if len(strings.Split(subdomain, ".")) < 3 {
		return ErrSubNotValid
	}

	subdomainsPool, err := c.projectProvider.GetOccupiedSubdomains(ctx)
	if err != nil {
		return err
	}

	for _, sN := range subdomainsPool {
		if subdomain == sN {
			return ErrSubAlrTaken
		}
	}

	domainsPool, err := c.projectProvider.GetDomainsPool(ctx)
	if err != nil {
		return err
	}

	sEls := strings.Split(subdomain, ".")
	domain := strings.Join(sEls[len(sEls)-2:], ".")
	for _, validDomain := range domainsPool {
		if validDomain == domain {
			return nil
		}
	}

	return ErrDomNotInPool
}

func (c *controller) endpointValidation(endpoint string) error {
	sEnd := strings.Split(endpoint, ":")
	if len(sEnd) != 2 {
		return ErrEpNotValid
	}

	if nil == net.ParseIP(sEnd[0]) {
		return ErrIpV4NotValid
	}

	port, err := strconv.Atoi(sEnd[1])
	if err != nil {
		return ErrPortNotValid
	}

	if port > 65_535 || port <= 0 {
		return ErrPortNotValid
	}

	return nil
}
