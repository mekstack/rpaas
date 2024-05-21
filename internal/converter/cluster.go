package xdsconverter

import (
	"strconv"
	"strings"
	xdsmodels "xds_server/internal/models"
)

func ClusterConverter(domainEp *xdsmodels.DomainEndpoint) xdsmodels.Cluster {
	var clusterName strings.Builder
	clusterName.WriteString(cluster)
	clusterName.WriteString("-")
	clusterName.WriteString(domainEp.Domain)

	port, _ := strconv.Atoi(strings.ReplaceAll(domainEp.Port, ":", ""))

	cl := xdsmodels.Cluster{
		Name: clusterName.String(),
		Endpoint: &xdsmodels.Endpoint{
			Address: domainEp.Host,
			Port:    uint32(port),
		},
	}

	return cl
}
