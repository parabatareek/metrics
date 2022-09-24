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
	// Инициализация канала
	channel := make(chan *metrics.Metrics)
	defer close(channel)

	// Инициализация структуры Metrics значениями runtime
	dataMetrics := metrics.NewMetrics()

	//for i := 0; i < 10; i++ {
	//	dataMetrics.Update()
	//}
	//fmt.Println(dataMetrics.PollCount)

	// Вызов обновления значений объекта Metrics в гоурутине.
	// Когда вы помещаете данные в канал, горутина блокируется до тех пор, пока данные не будут считаны
	// другой горутиной из этого канала.
	//https://habr.com/ru/post/490336/
	go runGetStats(dataMetrics, channel)

	runSendStats(channel)
}

// Обновление значений объекта Metrics
func runGetStats(dataMetrics *metrics.Metrics, channel chan *metrics.Metrics) {
	ticker := time.NewTicker(pollInterval)
	for {
		<-ticker.C
		dataMetrics.Update()
		channel <- dataMetrics
	}
}

func runSendStats(channel chan *metrics.Metrics) {
	ticker := time.NewTicker(reportInterval)
	var dataMetrics *metrics.Metrics
	for {
		<-ticker.C
		dataMetrics = <-channel
		fmt.Println(dataMetrics.PollCount)
	}
}
