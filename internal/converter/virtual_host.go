package xdsconverter

import (
	"strings"
	xdsmodels "xds_server/internal/models"
)

func VirtualHostConverter(domainEp *xdsmodels.DomainEndpoint) xdsmodels.VirtualHost {
	var clusterName strings.Builder
	clusterName.WriteString(cluster)
	clusterName.WriteString("-")
	clusterName.WriteString(domainEp.Domain)

	var routeName strings.Builder
	routeName.WriteString(route)
	routeName.WriteString("-")
	routeName.WriteString(domainEp.Domain)

	vh := xdsmodels.VirtualHost{
		Domain: domainEp.Domain,
		Route: &xdsmodels.Route{
			Name:        routeName.String(),
			Prefix:      "/",
			ClusterName: clusterName.String(),
		},
	}

	return vh
}
