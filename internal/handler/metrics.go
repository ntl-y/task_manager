package handler

import (
	"strconv"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	totalRequests = promauto.NewCounterVec(prometheus.CounterOpts{
		Namespace: "task_manager",
		Name:      "tasks_created_total",
		Help:      "The total number of handled HTTP requests.",
	}, []string{"status"})

	requestsDuration = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: "task_manager",
		Name:      "task_creation_duration_seconds",
		Help:      "A histogram of the HTTP request durations in seconds.",
		Buckets:   []float64{0.05, 0.1, 0.25, 0.5, 1, 2.5, 5, 10},
	}, []string{"status"})
)

func observeRequest(d time.Duration, status int) {
	totalRequests.WithLabelValues(strconv.Itoa(status)).Inc()
	requestsDuration.WithLabelValues(strconv.Itoa(status)).Observe(d.Seconds())
}
