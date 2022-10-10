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
	endpoint       = "http://127.0.0.1:8080"
	urlUpdate      = "/update/"
)

func main() {
	// Инициализация структуры Metrics значениями runtime
	dataMetrics := metrics.NewMetrics()

	// Вызов обновления значений объекта Metrics в гоурутине.
	go updStats(dataMetrics)

	// Отправка данных в гоурутине
	sendStats(dataMetrics)
}

// Обновление значений объекта Metrics
func updStats(dataMetrics *metrics.Metrics) {
	ticker := time.NewTicker(pollInterval)
	for {
		<-ticker.C
		dataMetrics.Update()
	}

}

func sendStats(datametrics *metrics.Metrics) {
	// Инициализация клиента
	client := &http.Client{}

	ticker := time.NewTicker(reportInterval)

	for {
		<-ticker.C

		// Формирование данных для отправки
		urlData := getParams(datametrics)
		for key, param := range urlData {
			urlParam := url.Values{}
			urlParam.Set(key, param)
			request := getRequest(&urlParam)
			getResponse(request, client)
		}
	}
}

// Формирование данных для отправки
func getParams(dataMetrics *metrics.Metrics) map[string]string {
	statType := reflect.TypeOf(dataMetrics).Elem()
	statVal := reflect.ValueOf(dataMetrics).Elem()

	urlData := make(map[string]string)

	for i := 0; i < statType.NumField(); i++ {
		fieldKind := statVal.Field(i).Kind()
		fieldName := statType.Field(i).Name
		fieldVal := statVal.Field(i)

		params := fmt.Sprintf("%v<%v>/<%s>/<%v>", urlUpdate, fieldKind, fieldName, fieldVal)

		urlData[fieldName] = params
	}
	return urlData
}

// Формирование request
func getRequest(urlData *url.Values) *http.Request {
	// Инициализация контекста
	//ctx, cancel := context.WithCancel(context.Background())
	//defer cancel()

	// Инициализация запроса
	request, err := http.NewRequestWithContext(context.Background(), http.MethodPost, endpoint, bytes.NewBufferString(urlData.Encode()))
	if err != nil {
		log.Fatal(err)
	}
	request.Header.Add("Content-Type", "text/plain")
	request.Header.Add("Content-Length", strconv.Itoa(len(urlData.Encode())))

	return request
}

// Отправка request
func getResponse(request *http.Request, client *http.Client) {
	// Отпавка данных
	response, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
}
