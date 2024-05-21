package serviceimpl

import (
	"github.com/envoyproxy/go-control-plane/pkg/cache/v3"
	"github.com/envoyproxy/go-control-plane/pkg/resource/v3"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"log/slog"
	"os"
	"testing"
	certclient "xds_server/internal/client/cert"
	xdsmodels "xds_server/internal/models"
)

func BenchmarkValidateEndpoint(b *testing.B) {
	domainEp := &xdsmodels.DomainEndpoint{
		Domain: "docs.test.ru",
		Host:   "127.0.0.1",
		Port:   ":2345",
	}

	ctrl := gomock.NewController(b)
	defer ctrl.Finish()

	mockedCertClient := certclient.NewMockClient(ctrl)

	cds, lds, rds := cache.NewLinearCache(resource.ClusterType), cache.NewLinearCache(resource.ListenerType), cache.NewLinearCache(resource.RouteType)

	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelError,
	}))

	impl := New(cds, lds, rds, mockedCertClient, logger, nodeId).(*service)

	b.ResetTimer()
	for i := 0; i < b.N; i += 1 {
		err := impl.validateEndpoint(domainEp)
		require.NoError(b, err)
	}
}
