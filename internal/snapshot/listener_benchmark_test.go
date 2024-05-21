package snapshot

import (
	"testing"
	xdsconverter "xds_server/internal/converter"
	xdsmodels "xds_server/internal/models"
)

const (
	certificateChain = "testCertificateChain"
	privateKey       = "testPrivateKey"
)

func BenchmarkGenerateListener(b *testing.B) {
	domainEp := &xdsmodels.DomainEndpoint{
		Domain: "docs.test.ru",
		Host:   "127.0.0.1",
		Port:   ":2345",
	}
	l := xdsconverter.ListenerConverter(domainEp, &xdsmodels.Cert{
		PrivateKey:       privateKey,
		CertificateChain: certificateChain,
	})

	b.ResetTimer()
	for i := 0; i < b.N; i += 1 {
		GenerateListener(l)
	}
}
