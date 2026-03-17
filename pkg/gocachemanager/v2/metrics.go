package v2

import (
	"sync/atomic"

	"github.com/prometheus/client_golang/prometheus"
)

const (
	defaultMetricsNamespace = "cache"
	storeLabelChain         = "chain"
)

// prometheusMetrics mirrors eko/gocache metrics: GaugeVec with labels service, store, metric.
type prometheusMetrics struct {
	collector *prometheus.GaugeVec
	service   string
	store     string

	hitCount      atomic.Uint64
	missCount     atomic.Uint64
	setSuccess    atomic.Uint64
	setError      atomic.Uint64
	deleteSuccess atomic.Uint64
	deleteError   atomic.Uint64
}

func newPrometheusMetrics(prefix, namespace string) *prometheusMetrics {
	if namespace == "" {
		namespace = defaultMetricsNamespace
	}

	m := &prometheusMetrics{
		service: prefix,
		store:   storeLabelChain,
		collector: prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Namespace: namespace,
				Name:      "collector",
				Help:      "Number of items / operations in cache (compatible with eko/gocache metrics)",
			},
			[]string{"service", "store", "metric"},
		),
	}

	prometheus.DefaultRegisterer.MustRegister(m.collector)

	return m
}

func (m *prometheusMetrics) recordHit() {
	v := m.hitCount.Add(1)
	m.collector.WithLabelValues(m.service, m.store, "hit_count").Set(float64(v))
}

func (m *prometheusMetrics) recordMiss() {
	v := m.missCount.Add(1)
	m.collector.WithLabelValues(m.service, m.store, "miss_count").Set(float64(v))
}

func (m *prometheusMetrics) recordSetSuccess() {
	v := m.setSuccess.Add(1)
	m.collector.WithLabelValues(m.service, m.store, "set_success").Set(float64(v))
}

func (m *prometheusMetrics) recordSetError() {
	v := m.setError.Add(1)
	m.collector.WithLabelValues(m.service, m.store, "set_error").Set(float64(v))
}

func (m *prometheusMetrics) recordDeleteSuccess() {
	v := m.deleteSuccess.Add(1)
	m.collector.WithLabelValues(m.service, m.store, "delete_success").Set(float64(v))
}

func (m *prometheusMetrics) recordDeleteError() {
	v := m.deleteError.Add(1)
	m.collector.WithLabelValues(m.service, m.store, "delete_error").Set(float64(v))
}
