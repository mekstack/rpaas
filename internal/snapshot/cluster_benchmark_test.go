package snapshot

import (
	"testing"
	xdsconverter "xds_server/internal/converter"
	xdsmodels "xds_server/internal/models"
)

func BenchmarkGenerateCluster(b *testing.B) {
	domainEp := &xdsmodels.DomainEndpoint{
		Domain: "docs.test.ru",
		Host:   "127.0.0.1",
		Port:   ":2345",
	}
	cl := xdsconverter.ClusterConverter(domainEp)

	b.ResetTimer()
	for i := 0; i < b.N; i += 1 {
		GenerateCluster(cl)
	}
}
