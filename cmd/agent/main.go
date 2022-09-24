package main

import (
	"github.com/parabatareek/metrics.git/internal/metrics"
	"time"
)

const (
	pollInterval   = 2 * time.Second
	reportInterval = 10 * time.Second
)

func main() {
	metrics := metrics.NewMetrics()
	go runGetStats(metrics)

	//fmt.Println(metrics)
	//fmt.Println(metrics)
}

func runGetStats(metrics *metrics.Metrics) {
	ticker := time.NewTicker(pollInterval)
	for {
		<-ticker.C
		metrics.Update()
	}
}
