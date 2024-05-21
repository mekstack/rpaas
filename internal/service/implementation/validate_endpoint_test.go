package serviceimpl

import (
	"github.com/envoyproxy/go-control-plane/pkg/cache/v3"
	"github.com/envoyproxy/go-control-plane/pkg/resource/v3"
	"github.com/gookit/goutil/testutil/assert"
	"go.uber.org/mock/gomock"
	"log/slog"
	"os"
	"testing"
	certclient "xds_server/internal/client/cert"
	xdsmodels "xds_server/internal/models"
	"xds_server/internal/service/errors"
)

func TestInspectEndpoint(t *testing.T) {
	cases := []*xdsmodels.DomainEndpoint{
		{
			Host:   "127.0.0.1",
			Domain: "docs.test.ru",
			Port:   ":2345",
		},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockedCertClient := certclient.NewMockClient(ctrl)

	cds, lds, rds := cache.NewLinearCache(resource.ClusterType), cache.NewLinearCache(resource.ListenerType), cache.NewLinearCache(resource.RouteType)

	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelError,
	}))

	impl := New(cds, lds, rds, mockedCertClient, logger, nodeId).(*service)

	for _, tCase := range cases {
		err := impl.validateEndpoint(tCase)
		assert.NoError(t, err)
	}
}

func TestInspectEndpointError(t *testing.T) {
	cases := []struct {
		name     string
		expErr   error
		domainEp *xdsmodels.DomainEndpoint
	}{
		{
			name:   "invalid domain",
			expErr: serviceerror.InvalidDomainError,
			domainEp: &xdsmodels.DomainEndpoint{
				Domain: "docs2222-?qte/qgeru",
				Host:   "127.0.0.1",
				Port:   ":2345",
			},
		},
		{
			name:   "invalid port",
			expErr: serviceerror.InvalidPortError,
			domainEp: &xdsmodels.DomainEndpoint{
				Domain: "docs.test.ru",
				Host:   "127.0.0.1",
				Port:   "-100",
			},
		},
		{
			name:   "invalid port",
			expErr: serviceerror.InvalidPortError,
			domainEp: &xdsmodels.DomainEndpoint{
				Domain: "docs.test.ru",
				Host:   "127.0.0.1",
				Port:   "2345",
			},
		},
		{
			name:   "invalid host",
			expErr: serviceerror.InvalidHostError,
			domainEp: &xdsmodels.DomainEndpoint{
				Domain: "docs.test.ru",
				Host:   "invalid_host",
				Port:   ":2345",
			},
		},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockedCertClient := certclient.NewMockClient(ctrl)

	cds, lds, rds := cache.NewLinearCache(resource.ClusterType), cache.NewLinearCache(resource.ListenerType), cache.NewLinearCache(resource.RouteType)

	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelError,
	}))

	impl := New(cds, lds, rds, mockedCertClient, logger, nodeId).(*service)

	for _, tCase := range cases {
		t.Run(tCase.name, func(t *testing.T) {
			err := impl.validateEndpoint(tCase.domainEp)
			assert.ErrIs(t, err, tCase.expErr)
		})
	}
}
