package snapshot

import (
	cluster "github.com/envoyproxy/go-control-plane/envoy/config/cluster/v3"
	endpoint "github.com/envoyproxy/go-control-plane/envoy/config/endpoint/v3"
	"github.com/envoyproxy/go-control-plane/pkg/cache/types"
	"google.golang.org/protobuf/types/known/durationpb"
	"time"
	xdsmodels "xds_server/internal/models"
)

import (
	core "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
)

func makeCluster(cl xdsmodels.Cluster) *cluster.Cluster {
	return &cluster.Cluster{
		Name:                 cl.Name,
		ConnectTimeout:       durationpb.New(5 * time.Second),
		ClusterDiscoveryType: &cluster.Cluster_Type{Type: cluster.Cluster_STATIC},
		LbPolicy:             cluster.Cluster_ROUND_ROBIN,
		LoadAssignment:       makeEndpoint(cl.Name, cl.Endpoint),
		DnsLookupFamily:      cluster.Cluster_V4_ONLY,
	}
}

func makeEndpoint(clusterName string, ep *xdsmodels.Endpoint) *endpoint.ClusterLoadAssignment {
	lbEndpoint := &endpoint.LbEndpoint{
		HostIdentifier: &endpoint.LbEndpoint_Endpoint{
			Endpoint: &endpoint.Endpoint{
				Address: &core.Address{
					Address: &core.Address_SocketAddress{
						SocketAddress: &core.SocketAddress{
							Protocol: core.SocketAddress_TCP,
							Address:  ep.Address,
							PortSpecifier: &core.SocketAddress_PortValue{
								PortValue: ep.Port,
							},
						},
					},
				},
			},
		},
		HealthStatus: core.HealthStatus_HEALTHY,
	}
	return &endpoint.ClusterLoadAssignment{
		ClusterName: clusterName,
		Endpoints: []*endpoint.LocalityLbEndpoints{{
			LbEndpoints: []*endpoint.LbEndpoint{lbEndpoint},
		}},
	}
}

func GenerateCluster(cl xdsmodels.Cluster) (m map[string]types.Resource, key string, value types.Resource) {
	m = make(map[string]types.Resource)
	v := makeCluster(cl)
	m[cl.Name] = v

	return m, cl.Name, v
}
