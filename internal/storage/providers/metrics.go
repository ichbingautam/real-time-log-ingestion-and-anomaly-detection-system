package storage

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

type Metrics struct {
	opsCounter   *prometheus.CounterVec
	opsDuration  *prometheus.HistogramVec
	bytesCounter prometheus.Counter
}

func NewStorageMetrics(reg prometheus.Registerer) *Metrics {
	m := &Metrics{
		opsCounter: promauto.With(reg).NewCounterVec(prometheus.CounterOpts{
			Name: "storage_operations_total",
			Help: "Total storage operations",
		}, []string{"type", "status"}),

		opsDuration: promauto.With(reg).NewHistogramVec(prometheus.HistogramOpts{
			Name:    "storage_operation_duration_seconds",
			Help:    "Storage operation duration distribution",
			Buckets: prometheus.DefBuckets,
		}, []string{"type"}),

		bytesCounter: promauto.With(reg).NewCounter(prometheus.CounterOpts{
			Name: "storage_bytes_processed_total",
			Help: "Total bytes processed by storage",
		}),
	}

	return m
}

func (m *Metrics) WrapRepository(repo Repository) Repository {
	return &instrumentedRepository{
		repo:    repo,
		metrics: m,
	}
}

type instrumentedRepository struct {
	repo    Repository
	metrics *Metrics
}

func (r *instrumentedRepository) Store(ctx context.Context, data []byte) error {
	start := time.Now()
	err := r.repo.Store(ctx, data)

	r.metrics.opsCounter.WithLabelValues("store", statusLabel(err)).Inc()
	r.metrics.opsDuration.WithLabelValues("store").Observe(time.Since(start).Seconds())
	r.metrics.bytesCounter.Add(float64(len(data)))

	return err
}

func statusLabel(err error) string {
	if err != nil {
		return "error"
	}
	return "success"
}