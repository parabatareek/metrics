package main

import (
	"bytes"
	"context"
	"fmt"
	"github.com/parabatareek/metrics.git/internal/metrics"
	"io"
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
	endpoint       = "http://127.0.0.1:8080/update/"
	//urlUpdate      = "/update/"
)

func main() {
	// Инициализация структуры Metrics значениями runtime
	dataMetrics := metrics.NewMetrics()

	// Вызов обновления значений объекта Metrics в гоурутине.
	go updStats(dataMetrics)

	// Отправка данных в гоурутине
	go sendStats(dataMetrics)
}

// Обновление значений объекта Metrics
func updStats(dataMetrics *metrics.Metrics) {
	ticker := time.NewTicker(pollInterval)
	<-ticker.C

	dataMetrics.Update()
	updStats(dataMetrics)
}

func sendStats(datametrics *metrics.Metrics) {
	ticker := time.NewTicker(reportInterval)
	<-ticker.C

	// Формирование данных для отправки
	urlData := getParams(datametrics)

	// Формирование request
	request := getRequest(urlData)

	getResponse(request)
	sendStats(datametrics)
}

// Формирование данных для отправки
func getParams(dataMetrics *metrics.Metrics) *url.Values {
	statType := reflect.TypeOf(dataMetrics).Elem()
	statVal := reflect.ValueOf(dataMetrics).Elem()

	urlData := url.Values{}

	for i := 0; i < statType.NumField(); i++ {
		fieldKind := statVal.Field(i).Kind()
		fieldName := statType.Field(i).Name
		fieldVal := statVal.Field(i)

		params := fmt.Sprintf("<%v>/<%s>/<%v>", fieldKind, fieldName, fieldVal)

		urlData.Set(fieldName, params)
	}
	return &urlData
}

func getRequest(urlData *url.Values) *http.Request {
	// Инициализация контекста
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	// Инициализация запроса
	request, err := http.NewRequestWithContext(ctx, http.MethodPost, endpoint, bytes.NewBufferString(urlData.Encode()))
	if err != nil {
		log.Fatal(err)
	}
	request.Header.Add("Content-Type", "text/plain")
	request.Header.Add("Content-Length", strconv.Itoa(len(urlData.Encode())))

	return request
}

func getResponse(request *http.Request) {
	// Инициализация клиента
	client := &http.Client{}

	// Отпавка данных
	response, err := client.Do(request)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(response.Body)

	if err != nil {
		log.Fatal(err)
	}
}
