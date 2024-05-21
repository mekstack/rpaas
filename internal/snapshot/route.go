package snapshot

import (
	route "github.com/envoyproxy/go-control-plane/envoy/config/route/v3"
	"github.com/envoyproxy/go-control-plane/pkg/cache/types"
	xdsmodels "xds_server/internal/models"
)

const (
	virtualHostName = "local_service"
)

func makeRoutes(r *xdsmodels.Route) []*route.Route {
	return []*route.Route{{
		Match: &route.RouteMatch{
			PathSpecifier: &route.RouteMatch_Prefix{
				Prefix: r.Prefix,
			},
		},
		Action: &route.Route_Route{
			Route: &route.RouteAction{
				ClusterSpecifier: &route.RouteAction_Cluster{
					Cluster: r.ClusterName,
				},
			},
		},
	}}
}

func makeRouteConfiguration(vh xdsmodels.VirtualHost) *route.RouteConfiguration {
	return &route.RouteConfiguration{
		Name: vh.Route.Name,
		VirtualHosts: []*route.VirtualHost{{
			Name:    virtualHostName,
			Domains: []string{vh.Domain},
			Routes:  makeRoutes(vh.Route),
		}},
	}
}

func GenerateRoute(vh xdsmodels.VirtualHost) (m map[string]types.Resource, key string, value types.Resource) {
	m = make(map[string]types.Resource)
	v := makeRouteConfiguration(vh)
	m[vh.Route.Name] = v

	return m, vh.Route.Name, v
}
