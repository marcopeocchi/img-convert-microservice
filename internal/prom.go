package internal

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	TimePerOpGauge = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "time_per_conversion_ns",
		Help: "Time to complete a conversion",
	})
	OpsCounter = promauto.NewCounter(prometheus.CounterOpts{
		Name: "processed_counter",
		Help: "Number of request succesfully processed",
	})
)
