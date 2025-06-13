// Package prometheus implements a metric collector [backend.MetricCollector] client interface for the prometheus metric collector.
package prometheus

import (
	"context"

	"github.com/blueprint-uservices/blueprint/runtime/core/backend"
	"go.opentelemetry.io/otel"
	prometheus_exporter "go.opentelemetry.io/otel/exporters/prometheus"
	"go.opentelemetry.io/otel/metric"
	metricsdk "go.opentelemetry.io/otel/sdk/metric"
)

// PrometheusMetricCollector implements the runtime backend instance that implements the backend/metric.MetricCollector interface.
type PrometheusMetricCollector struct {
	mp *metricsdk.MeterProvider
}

func (p *PrometheusMetricCollector) GetMetricProvider(ctx context.Context) (metric.MeterProvider, error) {
	return p.mp, nil
}

// Returns a new instance of PrometheusMetricCollector.
// Configures opentelemetry to export metrics to a Prometheus server hosted at address `addr`.
func NewPrometheusMetricCollector(ctx context.Context, addr string) (*PrometheusMetricCollector, error) {
	exp, err := prometheus_exporter.New()
	if err != nil {
		return nil, err
	}
	mp := metricsdk.NewMeterProvider(
		metricsdk.WithReader(exp),
	)
	otel.SetMeterProvider(mp)
	mc := &PrometheusMetricCollector{mp}
	backend.SetDefaultMetricCollector(mc)
	return mc, nil
}
