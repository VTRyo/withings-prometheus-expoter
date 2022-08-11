package main

import (
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	weightGage = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: "",
			Name:      "",
			Help:      "",
		}, []string{""},
	)
)

func main() {
	prometheus.Register()

	http.Handle("/metrics", promhttp.Handler())
	go func() {
		setValue()
	}()
	log.Fatal(http.ListenAndServe("8080", nil))
}

func setValue() {
	// 1日間隔で情報を取得する
	// Labelを付与する
	// 値をSetする
}
