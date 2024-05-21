package serviceimpl

import (
	"context"
	"github.com/envoyproxy/go-control-plane/pkg/cache/v3"
	"github.com/envoyproxy/go-control-plane/pkg/resource/v3"
	"github.com/stretchr/testify/assert"
	"go.uber.org/goleak"
	"go.uber.org/mock/gomock"
	"log/slog"
	"os"
	"testing"
	certclient "xds_server/internal/client/cert"
	xdsmodels "xds_server/internal/models"
)

const (
	certificateChain = "testCertificateChain"
	privateKey       = "testPrivateKey"
	nodeId           = "0"
)

func TestApplyEndpoint(t *testing.T) {
	//GOROUTINES LEAK DETECTOR
	defer goleak.VerifyNone(t)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	mockedCertClient := certclient.NewMockClient(ctrl)
	mockedCertClient.EXPECT().Cert(ctx, "docs.test.ru").Return(&xdsmodels.Cert{CertificateChain: certificateChain, PrivateKey: privateKey}, nil)

	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelError,
	}))

	cdsCache, ldsCache, rdsCache := cache.NewLinearCache(resource.ClusterType), cache.NewLinearCache(resource.ListenerType), cache.NewLinearCache(resource.RouteType)
	i := New(cdsCache, ldsCache, rdsCache, mockedCertClient, logger, nodeId)

	cases := []*xdsmodels.DomainEndpoint{
		{
			Host:   "127.0.0.1",
			Domain: "docs.test.ru",
			Port:   ":2345",
		},
	}

	for _, tCase := range cases {
		err := i.ApplyEndpoint(context.Background(), tCase)
		assert.NoError(t, err)
	}
}
