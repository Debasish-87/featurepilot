package metrics

import "github.com/prometheus/client_golang/prometheus"

var (
	EvaluationTotal = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "evaluation_total",
			Help: "Total feature evaluations",
		},
	)

	CacheHitTotal = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "cache_hit_total",
			Help: "Total cache hits",
		},
	)

	CacheMissTotal = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "cache_miss_total",
			Help: "Total cache misses",
		},
	)

	EvaluationErrorsTotal = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "evaluation_errors_total",
			Help: "Total evaluation errors",
		},
	)

	EvaluationDuration = prometheus.NewHistogram(
		prometheus.HistogramOpts{
			Name:    "evaluation_duration_seconds",
			Help:    "Evaluation latency",
			Buckets: prometheus.DefBuckets,
		},
	)
)

func Register() {
	prometheus.MustRegister(
		EvaluationTotal,
		CacheHitTotal,
		CacheMissTotal,
		EvaluationErrorsTotal,
		EvaluationDuration,
	)
}