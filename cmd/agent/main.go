package main

import (
	"fmt"
	"github.com/parabatareek/metrics.git/internal/metrics"
	"time"
)

const (
	pollInterval   = 2 * time.Second
	reportInterval = 10 * time.Second
)

func main() {
	metrics := metrics.NewMetrics()
	fmt.Println(metrics)
}
