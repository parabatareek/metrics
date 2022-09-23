package metrics

import (
	"math/rand"
	"reflect"
	"runtime"
	"time"
)

type Metrics struct {
	Alloc         float64
	BuckHashSys   float64
	Frees         float64
	GCCPUFraction float64
	GCSys         float64
	HeapAlloc     float64
	HeapIdle      float64
	HeapInuse     float64
	HeapObjects   float64
	HeapReleased  float64
	HeapSys       float64
	LastGC        float64
	Lookups       float64
	MCacheInuse   float64
	MCacheSys     float64
	MSpanInuse    float64
	MSpanSys      float64
	Mallocs       float64
	NextGC        float64
	NumForcedGC   float64
	NumGC         float64
	OtherSys      float64
	PauseTotalNs  float64
	StackInuse    float64
	StackSys      float64
	Sys           float64
	TotalAlloc    float64
	RandomValue   float64
	PollCount     int64
}

// Конструктор Metrics
func NewMetrics() *Metrics {
	var metrics Metrics
	var rtm runtime.MemStats
	runtime.ReadMemStats(&rtm)

	// Рефлексия runtime
	rtmType := reflect.TypeOf(rtm)
	rtmVal := reflect.ValueOf(rtm)

	// Рефлексия Metrics
	metType := reflect.TypeOf(metrics)
	// Разыменование структуры Metrics для доступа к записи значений в структуру
	// https://golang-blog.blogspot.com/2020/06/laws-of-reflection-in-golang.html
	metVal := reflect.ValueOf(&metrics).Elem()

	// Запись значений runtime в структуру Metrics
	for i := 0; i < rtmType.NumField(); i++ {
		rtmFieldName := rtmType.Field(i).Name
		metFieldName := metType.Field(i).Name

		if rtmFieldName == metFieldName {
			field := metVal.FieldByName(metFieldName)
			newVal := getValue(rtmVal.Field(i))
			field.SetFloat(newVal)
		}
	}
	rand.Seed(time.Now().UnixNano())
	metrics.RandomValue = rand.Float64()
	return &metrics
}

func getValue(fieldVal reflect.Value) (val float64) {
	switch fieldVal.Kind() {
	case reflect.Uint32, reflect.Uint64:
		val = float64(fieldVal.Uint())
	case reflect.Float32, reflect.Float64:
		val = fieldVal.Float()
	}
	return
}
