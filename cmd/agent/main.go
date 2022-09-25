package main

import (
	"bytes"
	"context"
	"fmt"
	"github.com/parabatareek/metrics.git/internal/metrics"
	"log"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"time"
)

const (
	pollInterval   = 2 * time.Second
	reportInterval = 10 * time.Second
	endpoint       = "127.0.0.1:8080"
)

func main() {
	// Инициализация канала
	getChannel := make(chan *metrics.Metrics)
	defer close(getChannel)

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
	go runGetStats(dataMetrics, getChannel)

	runReadStats(getChannel)
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

// Чтение значений обновленного объекта Metrics
func runReadStats(channel chan *metrics.Metrics) {
	ticker := time.NewTicker(reportInterval)
	var dataMetrics *metrics.Metrics
	for {
		<-ticker.C
		dataMetrics = <-channel
		runSendStats(dataMetrics)
	}
}

// Подготовка и отправка данных на сервер
func runSendStats(dataMetrics *metrics.Metrics) {
	statType := reflect.TypeOf(dataMetrics).Elem()
	statVal := reflect.ValueOf(dataMetrics).Elem()

	data := url.Values{}
	urlStr := "/update/"

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	for i := 0; i < statType.NumField(); i++ {
		fieldKind := statVal.Field(i).Kind()
		fieldName := statType.Field(i).Name
		fieldVal := statVal.Field(i)
		//fieldVal := getStrValue(statVal.Field(i))

		params := fmt.Sprintf("<%v>/<%s>/<%v>", fieldKind, fieldName, fieldVal)
		urlStr += params

		data.Set("url", urlStr)

		request, err := http.NewRequestWithContext(ctx, http.MethodPost, endpoint, bytes.NewBufferString(data.Encode()))
		if err != nil {
			log.Fatal(err)
		}

		request.Header.Add("Content-Type", "text/plain")
		request.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))
	}
}

// Приведение типов reflect.Value к float64
func getStrValue(fieldVal reflect.Value) (val string) {
	switch fieldVal.Kind() {
	case reflect.Uint32, reflect.Uint64:
		val = string(fieldVal.Uint())
	case reflect.Float32, reflect.Float64:
		val = strconv.FormatFloat(fieldVal.Float(), 'f', -1, 64)
	}
	return
}
