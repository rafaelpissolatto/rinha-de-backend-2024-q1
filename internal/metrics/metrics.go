/* Responsible to collect metrics from the system and send them to the server. Using OpenTelemetry */

package metrics

import (
	"errors"
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Monitor struct {
	Counter *prometheus.CounterVec
	Gauge   *prometheus.GaugeVec
}

func serveMetrics() {
	// http.Handle("/metrics", promhttp.Handler())
	http.Handle("/metrics", promhttp.Handler())
	err := http.ListenAndServe(":8091", nil)
	if err != nil {
		log.Println("[ERROR] error serving metrics", err)
		return
	}
}

// InitMetrics starts the metrics server
func InitMetrics() {
	go serveMetrics()
}

// NewMonitor creates a new monitor
func NewMonitor() (*Monitor, error) {
	counter := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests made.",
		},
		[]string{"code", "method"},
	)

	gauge := prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "http_requests_in_progress",
			Help: "Number of HTTP requests currently in progress.",
		},
		[]string{"method"},
	)

	if err := prometheus.Register(counter); err != nil {
		return nil, errors.New("could not register counter")
	}

	if err := prometheus.Register(gauge); err != nil {
		return nil, errors.New("could not register gauge")
	}

	return &Monitor{
		Counter: counter,
		Gauge:   gauge,
	}, nil
}
