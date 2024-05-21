package snapshot

import (
	"testing"
	xdsconverter "xds_server/internal/converter"
	xdsmodels "xds_server/internal/models"
)

func BenchmarkGenerateRoute(b *testing.B) {
	domainEp := &xdsmodels.DomainEndpoint{
		Domain: "docs.test.ru",
		Host:   "127.0.0.1",
		Port:   ":2345",
	}

	vh := xdsconverter.VirtualHostConverter(domainEp)

	b.ResetTimer()
	for i := 0; i < b.N; i += 1 {
		GenerateRoute(vh)
	}
}
