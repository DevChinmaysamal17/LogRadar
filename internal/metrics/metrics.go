package metrics

import (
	"log"
	"net/http"

	//Required prometheus imports
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (

	//
	EventsProcessed = promauto.NewCounter(
		prometheus.CounterOpts{
			Name: "logradar_events_processed_total", //Incremented in main.go after every log entry
			Help: "Total number of successfully parsed log events.",
		},
	)
	AlertsTriggered = promauto.NewCounter(
		prometheus.CounterOpts{
			Name: "logradar_alerts_triggered_total", // Incremented in workers.go after every error log entry
			Help: "Total number of alerts triggered.",
		},
	)
)

func Serve(port string) {

	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(":"+port, nil))

}
