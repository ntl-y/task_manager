package metrics

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
)

func Listen(addr string) error {
	logrus.Infof("Metrics server starting at %s", addr)
	mux := http.NewServeMux()
	mux.Handle("/metrics", promhttp.Handler())

	return http.ListenAndServe(addr, mux)
}
