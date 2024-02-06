package project_controller

import (
	"context"
	"fmt"
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
		// TODO: Handle potential error
		sRoute := strings.Split(route, ":")
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
	// TODO: Add validator that subdomain is free and there is domain form this subdomain in domains pool
	subdomain := route.Subdomain.Name

	endpoint := route.Endpoint

	// TODO: Handle potential error
	epEls := strings.Split(endpoint, ":")

	if nil == net.ParseIP(epEls[0]) {
		return nil, fmt.Errorf("Invalid IPv4 address")
	}

	port, err := strconv.Atoi(epEls[1])
	if err != nil {
		return nil, err
	}

	if port > 65_535 || port <= 0 {
		return nil, fmt.Errorf("Invalid Port")
	}

	if err := c.projectProvider.AddRouteToProject(ctx, projectCode, endpoint, subdomain); err != nil {
		return nil, err
	}

	return c.GetProject(ctx, projectCode)
}
