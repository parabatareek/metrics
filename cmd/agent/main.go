package main

import (
	"fmt"
	"github.com/parabatareek/metrics.git/internal/metrics"
)

func main() {
	Metrics := metrics.NewMetrics()
	fmt.Println(Metrics)
}
