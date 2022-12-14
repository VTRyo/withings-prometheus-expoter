package main

import (
	"log"
	"net/http"
	"time"

	"github.com/VTRyo/withings-exporter/withings"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	weightGauge = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: "withings",
			Name:      "weight",
			Help:      "Weight value",
		}, []string{"device_id"},
	)

	fatGauge = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: "withings",
			Name:      "fat_rate",
			Help:      "Fat Rate",
		}, []string{"device_id"},
	)
)

func main() {
	prometheus.Register(weightGauge)
	prometheus.Register(fatGauge)

	http.Handle("/metrics", promhttp.Handler())
	go func() {
		setValue()
	}()
	log.Fatal(http.ListenAndServe(":8181", nil))
}

func setValue() {
	withings.SetToken()
	weight := withings.GetWeight().Value
	weightLabels := prometheus.Labels{
		"device_id": withings.GetWeight().DeviceID,
	}

	fat := withings.GetFat().Value
	fatLabels := prometheus.Labels{
		"device_id": withings.GetFat().DeviceID,
	}

	weightGauge.With(weightLabels).Set(weight)
	fatGauge.With(fatLabels).Set(fat)
	time.Sleep(86400 * time.Second) // 1日間隔
}
