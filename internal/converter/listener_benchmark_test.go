package xdsconverter

import (
	"testing"
	xdsmodels "xds_server/internal/models"
)

const (
	CertificateChain = "testCertificateChain"
	PrivateKey       = "testPrivateKey"
)

func BenchmarkListenerConverter(b *testing.B) {
	domainEp := &xdsmodels.DomainEndpoint{
		Host:   "127.0.0.1",
		Domain: "docs.test.ru",
		Port:   ":2345",
	}

	cert := &xdsmodels.Cert{
		CertificateChain: CertificateChain,
		PrivateKey:       PrivateKey,
	}

	b.ResetTimer()
	for i := 0; i < b.N; i += 1 {
		ListenerConverter(domainEp, nil)
		ListenerConverter(domainEp, cert)
	}

}
