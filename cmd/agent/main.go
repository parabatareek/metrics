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
	// Инициализация канала
	updParamChan := make(chan *metrics.Metrics)
	defer close(updParamChan)

	// Инициализация структуры Metrics значениями runtime
	dataMetrics := metrics.NewMetrics()

	// Вызов обновления значений объекта Metrics в гоурутине.
	// Когда вы помещаете данные в канал, горутина блокируется до тех пор, пока данные не будут считаны
	// другой горутиной из этого канала.
	//https://habr.com/ru/post/490336/
	go updStats(dataMetrics, updParamChan)

	// Чтение обновленных данных runtime в гоурутине
	go readStats(updParamChan)

	// Инициализация контекста
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	// Инициализация клиента
	client := &http.Client{}

	select {
	case stats := <-updParamChan:
		urlData := setParams(stats)
		sendData(ctx, urlData, client)
	}
}

func sendData(ctx context.Context, urlData *url.Values, client *http.Client) {
	request, err := http.NewRequestWithContext(ctx, http.MethodPost, endpoint, bytes.NewBufferString(urlData.Encode()))
	if err != nil {
		log.Fatal(err)
	}
	request.Header.Add("Content-Type", "text/plain")
	request.Header.Add("Content-Length", strconv.Itoa(len(urlData.Encode())))

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

// Обновление значений объекта Metrics
func updStats(dataMetrics *metrics.Metrics, channel chan *metrics.Metrics) {
	ticker := time.NewTicker(pollInterval)
	for {
		<-ticker.C
		dataMetrics.Update()
		channel <- dataMetrics
	}
}

// Чтение значений обновленного объекта Metrics
func readStats(channel chan *metrics.Metrics) {
	ticker := time.NewTicker(reportInterval)
	//var dataMetrics *metrics.Metrics
	for {
		<-ticker.C
		<-channel
	}
}

func setParams(dataMetrics *metrics.Metrics) *url.Values {
	statType := reflect.TypeOf(dataMetrics).Elem()
	statVal := reflect.ValueOf(dataMetrics).Elem()

	urlData := url.Values{}

	for i := 0; i < statType.NumField(); i++ {
		fieldKind := statVal.Field(i).Kind()
		fieldName := statType.Field(i).Name
		fieldVal := statVal.Field(i)

		params := fmt.Sprintf("<%v>/<%s>/<%v>", fieldKind, fieldName, fieldVal)
		//urlParams := urlUpdate + params

		urlData.Set(fieldName, params)
	}
	return &urlData
}
