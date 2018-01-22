package metrics

import (
	"github.com/go-kit/kit/metrics"
	"github.com/go-kit/kit/metrics/expvar"
)

var (
	// Metrics -
	Metrics struct {
		SearchTime        metrics.Counter
		AddressesSearched metrics.Counter
		AddressesMatched  metrics.Counter
	}
)

// Initialize -
func Initialize() {
	SearchTime = expvar.NewCounter("SearchTime")
	AddressesSearched = expvar.NewCounter("AddressesSearched")
	AddressesMatched = expvar.NewCounter("AddressesMatched")
}
