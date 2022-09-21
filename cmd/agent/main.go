package main

import (
	"fmt"
	"github.com/parabatareek/metrics.git/internal/metrics"
)

func main() {
	m := metrics.NewMetrics()
	fmt.Println(m)
}
