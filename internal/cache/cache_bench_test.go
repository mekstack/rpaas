package xdscache

import (
	"context"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"log/slog"
	"os"
	"testing"
	certclient "xds_server/internal/client/cert"
	allroutesclient "xds_server/internal/client/routes"
	xdsmodels "xds_server/internal/models"
)

const (
	certificateChain = "testCertificateChain"
	privateKey       = "testPrivateKey"
)

func BenchmarkInitializeCache(b *testing.B) {
	ctrl := gomock.NewController(b)
	defer ctrl.Finish()

	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelError,
	}))

	mockedCertClient := certclient.NewMockClient(ctrl)

	mockedRoutesClient := allroutesclient.NewMockClient(ctrl)

	b.ResetTimer()
	for i := 0; i < b.N; i += 1 {
		b.StopTimer()
		domainEpCh := make(chan *xdsmodels.DomainEndpoint, 5)
		domainEpCh <- &xdsmodels.DomainEndpoint{
			Domain: "docs.test.ru",
			Host:   "127.0.0.1",
			Port:   ":2345",
		}
		close(domainEpCh)

		mockedRoutesClient.EXPECT().Routes(gomock.Any()).Return(domainEpCh, nil)
		mockedCertClient.EXPECT().Cert(gomock.Any(), "docs.test.ru").Return(&xdsmodels.Cert{CertificateChain: certificateChain, PrivateKey: privateKey}, nil)
		b.StartTimer()

		_, _, _, err := InitializeCache(context.Background(), mockedCertClient, mockedRoutesClient, logger)
		require.NoError(b, err)
	}

}
