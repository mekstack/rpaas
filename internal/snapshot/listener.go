package snapshot

import (
	xdsmodels "xds_server/internal/models"

	listener "github.com/envoyproxy/go-control-plane/envoy/config/listener/v3"
	router "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/http/router/v3"
	hcm "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/network/http_connection_manager/v3"
	tlsv3 "github.com/envoyproxy/go-control-plane/envoy/extensions/transport_sockets/tls/v3"
	"github.com/envoyproxy/go-control-plane/pkg/cache/types"
	"github.com/envoyproxy/go-control-plane/pkg/resource/v3"
	"github.com/envoyproxy/go-control-plane/pkg/wellknown"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/wrapperspb"

	core "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
)

const (
	httpv1                 = "http/1.1"
	httpv2                 = "h2"
	managerStatPrefix      = "ingress_http"
	tlsTransportSocketName = "envoy.transport_sockets.tls"
	xdsCluster             = "xds_cluster"
)

func configSource() *core.ConfigSource {
	source := &core.ConfigSource{}
	source.ResourceApiVersion = resource.DefaultAPIVersion
	source.ConfigSourceSpecifier = &core.ConfigSource_ApiConfigSource{
		ApiConfigSource: &core.ApiConfigSource{
			TransportApiVersion:       resource.DefaultAPIVersion,
			ApiType:                   core.ApiConfigSource_DELTA_GRPC,
			SetNodeOnFirstMessageOnly: true,
			GrpcServices: []*core.GrpcService{{
				TargetSpecifier: &core.GrpcService_EnvoyGrpc_{
					EnvoyGrpc: &core.GrpcService_EnvoyGrpc{ClusterName: xdsCluster},
				},
			}},
		},
	}
	return source
}

func createDownstreamTlsContextAny(cert *xdsmodels.Cert) *anypb.Any {
	if cert == nil {
		return nil
	}

	downstreamTlsContext := &tlsv3.DownstreamTlsContext{
		CommonTlsContext: &tlsv3.CommonTlsContext{
			TlsCertificates: []*tlsv3.TlsCertificate{{
				CertificateChain: &core.DataSource{
					Specifier: &core.DataSource_InlineString{
						InlineString: cert.CertificateChain,
					},
				},
				PrivateKey: &core.DataSource{
					Specifier: &core.DataSource_InlineString{
						InlineString: cert.PrivateKey,
					},
				},
			}},
			AlpnProtocols: []string{httpv1, httpv2},
		},
	}

	a, _ := anypb.New(downstreamTlsContext)
	return a
}

func makeHTTPListener(l *xdsmodels.Listener) *listener.Listener {
	source := configSource()
	certSetting := createDownstreamTlsContextAny(l.Cert)

	routerConfig, err := anypb.New(&router.Router{})
	if err != nil {
		panic(err)
	}

	manager := &hcm.HttpConnectionManager{
		CodecType:  hcm.HttpConnectionManager_AUTO,
		StatPrefix: managerStatPrefix,
		RouteSpecifier: &hcm.HttpConnectionManager_Rds{
			Rds: &hcm.Rds{
				ConfigSource:    source,
				RouteConfigName: l.RouteConfigName,
			},
		},
		HttpFilters: []*hcm.HttpFilter{{
			Name:       wellknown.Router,
			ConfigType: &hcm.HttpFilter_TypedConfig{TypedConfig: routerConfig},
		}},
	}

	pbst, _ := anypb.New(manager)

	filterChain := &listener.FilterChain{
		Filters: []*listener.Filter{{
			Name: wellknown.HTTPConnectionManager,
			ConfigType: &listener.Filter_TypedConfig{
				TypedConfig: pbst,
			},
		}},
	}

	if certSetting != nil {
		filterChain.TransportSocket = &core.TransportSocket{
			Name: tlsTransportSocketName,
			ConfigType: &core.TransportSocket_TypedConfig{
				TypedConfig: certSetting,
			},
		}
	}

	return &listener.Listener{
		Name: l.Name,
		Address: &core.Address{
			Address: &core.Address_SocketAddress{
				SocketAddress: &core.SocketAddress{
					Protocol: core.SocketAddress_TCP,
					Address:  l.Address,
					PortSpecifier: &core.SocketAddress_PortValue{
						PortValue: l.Port,
					},
				},
			},
		},
		FilterChains:    []*listener.FilterChain{filterChain},
		EnableReusePort: wrapperspb.Bool(true),
	}
}

func GenerateListener(l *xdsmodels.Listener) (m map[string]types.Resource, key string, value types.Resource) {
	m = make(map[string]types.Resource)
	v := makeHTTPListener(l)
	m[l.Name] = v

	return m, l.Name, v
}
