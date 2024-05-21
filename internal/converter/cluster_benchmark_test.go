package xdsconverter

import (
	"testing"
	xdsmodels "xds_server/internal/models"
)

func BenchmarkClusterConverter(b *testing.B) {
	domainEp := &xdsmodels.DomainEndpoint{
		Host:   "127.0.0.1",
		Domain: "docs.test.ru",
		Port:   ":2345",
	}

	b.ResetTimer()
	for i := 0; i < b.N; i += 1 {
		ClusterConverter(domainEp)
	}
}
