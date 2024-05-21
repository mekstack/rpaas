package serviceimpl

import (
	"context"
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

func BenchmarkServiceApplyEndpoint(b *testing.B) {
	ctrl := gomock.NewController(b)
	defer ctrl.Finish()

	mockedCertClient := certclient.NewMockClient(ctrl)

	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelError,
	}))

	cds, lds, rds := cache.NewLinearCache(resource.ClusterType), cache.NewLinearCache(resource.ListenerType), cache.NewLinearCache(resource.RouteType)
	impl := New(cds, lds, rds, mockedCertClient, logger, nodeId)

	domainEp := &xdsmodels.DomainEndpoint{
		Host:   "127.0.0.1",
		Domain: "docs.test.ru",
		Port:   ":2345",
	}

	b.ResetTimer()
	for i := 0; i < b.N; i += 1 {
		b.StopTimer()
		mockedCertClient.EXPECT().Cert(gomock.Any(), "docs.test.ru").Return(&xdsmodels.Cert{PrivateKey: privateKey, CertificateChain: certificateChain}, nil)
		b.StartTimer()

		err := impl.ApplyEndpoint(context.Background(), domainEp)
		require.NoError(b, err)
	}
}
