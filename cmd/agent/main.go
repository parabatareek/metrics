package main

import (
	"fmt"
	"github.com/parabatareek/metrics.git/internal/metrics"
)

func main() {
	metrics := metrics.NewMetrics()
	fmt.Println(metrics)
	metrics.Update()
}
