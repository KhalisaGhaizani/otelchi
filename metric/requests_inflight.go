package metric

import (
	"fmt"
	"net/http"

	"go.opentelemetry.io/otel/attribute"
	otelmetric "go.opentelemetry.io/otel/metric"
)

const (
	metricNameRequestInFlight = "requests_inflight"
	metricUnitRequestInFlight = "{count}"
	metricDescRequestInFlight = "Measures the number of requests currently being processed by the server."
)

// [RequestInFlight] is a metrics recorder for recording the number of requests in flight.
func NewRequestInFlight(cfg BaseConfig) func(next http.Handler) http.Handler {
	// init metric, here we are using counter for capturing request in flight
	counter, err := cfg.Meter.Int64UpDownCounter(
		metricNameRequestInFlight,
		otelmetric.WithDescription(metricDescRequestInFlight),
		otelmetric.WithUnit(metricUnitRequestInFlight),
	)
	if err != nil {
		panic(fmt.Sprintf("unable to create %s counter: %v", metricNameRequestInFlight, err))
	}

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// build attributes without status_code or outcome yet
			baseAttrs := cfg.AttributesFunc(r)

			// get recording response writer
			rrw := getRRW(w)
			defer putRRW(rrw)

			// increment inflight
			counter.Add(r.Context(), 1,
				otelmetric.WithAttributes(baseAttrs...),
			)

			next.ServeHTTP(rrw.writer, r)

			// determine success/failure
			outcome := getOutcome(rrw.statusCode)

			// final attributes (with status_code and outcome)
			finalAttrs := append(
				baseAttrs,
				attribute.Int("http.status_code", rrw.statusCode),
				attribute.String("http.response_outcome", outcome),
			)

			// decrement inflight
			counter.Add(r.Context(), -1,
				otelmetric.WithAttributes(finalAttrs...),
			)
		})
	}
}
