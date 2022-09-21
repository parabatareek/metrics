package metrics

import (
	"math/rand"
	"reflect"
	"runtime"
	"time"
)

const (
	pollInterval   = 2 * time.Second
	reportInterval = 10 * time.Second
)

//func GetMetricsName() []string {
//	return []string{"Alloc", "BuckHashSys", "Frees", "GCCPUFraction", "GCSys", "HeapAlloc", "HeapIdle", "HeapInuse", "HeapObjects", "HeapReleased", "HeapSys", "LastGC", "Lookups", "MCacheInuse", "MCacheSys", "MSpanInuse", "MSpanSys", "Mallocs", "NextGC", "NumForcedGC", "NumGC", "OtherSys", "PauseTotalNs", "StackInuse", "StackSys", "Sys", "TotalAlloc"}
//}
//
//func GetMetrics() (metrics map[string]float64) {
//	rtm := runtime.MemStats{}
//	runtime.ReadMemStats(&rtm)
//	statsVal := reflect.ValueOf(rtm)
//	statsType := reflect.TypeOf(rtm)
//
//	for _, val := range GetMetricsName() {
//		for i := 0; i < statsVal.NumField(); i++ {
//			if val == statsType.Field(i).Name {
//				metrics[val] = statsVal.Field(i).Float()
//			}
//		}
//	}
//	return metrics
//}

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

// Конструктор Metrics{}
func NewMetrics() *Metrics {
	var rtm runtime.MemStats
	runtime.ReadMemStats(&rtm)
	m := *new(Metrics)
	metType := reflect.TypeOf(m)
	metVal := reflect.ValueOf(m)
	rtmType := reflect.TypeOf(rtm)
	rtmVal := reflect.ValueOf(rtm)

	// Запись текущих значений runtime в структуру
	for i := 0; i < rtmType.NumField(); i++ {

		if metType.Field(i).Name == rtmType.Field(i).Name {
			switch rtmVal.Field(i).Kind() {
			case reflect.Uint:
				v := float64(rtmVal.Field(i).Uint())
				metVal.Field(i).SetFloat(v)
			}
		}
		//fieldName := rtmType.Field(i).Name
		//metVal.FieldByName(rtmType.Field(i).Name).SetFloat(rtmVal.Field(i).Float())
	}

	// Случайное значение для m.RandomValue
	rand.Seed(time.Now().UnixNano())
	m.RandomValue = rand.Float64()

	return &m
}

// Запись данных runtime в Metrics
func (m *Metrics) Update() {
	var rtm *runtime.MemStats

	metType := reflect.TypeOf(&m)
	metVal := reflect.ValueOf(&m)
	rtmType := reflect.TypeOf(rtm)
	rtmVal := reflect.ValueOf(rtm)

	for {
		runtime.ReadMemStats(rtm)

		time.After(pollInterval)
		for i := 0; i < metType.NumField(); i++ {
			metVal.FieldByName(rtmType.Field(i).Name).SetFloat(rtmVal.Field(i).Float())
		}

		// Случайное значение для m.RandomValue
		rand.Seed(time.Now().UnixNano())
		m.RandomValue = rand.Float64()

		// Увеличиваем счетчик запросов к runtime
		m.PollCount += 1
	}
}

//
//func NewMetrics() *Metrics {
//	var rtm runtime.MemStats
//
//	runtime.ReadMemStats(&rtm)
//
//	return &Metrics{
//		Alloc:         (rtm.Alloc),
//		BuckHashSys:   (rtm.BuckHashSys),
//		Frees:         (rtm.Frees),
//		GCCPUFraction: rtm.GCCPUFraction,
//		GCSys:         (rtm.GCSys),
//		HeapAlloc:     (rtm.HeapAlloc),
//		HeapIdle:      (rtm.HeapIdle),
//		HeapInuse:     (rtm.HeapInuse),
//		HeapObjects:   (rtm.HeapObjects),
//		HeapReleased:  (rtm.HeapReleased),
//		HeapSys:       (rtm.HeapSys),
//		LastGC:        (rtm.LastGC),
//		Lookups:       (rtm.Lookups),
//		MCacheInuse:   (rtm.MCacheInuse),
//		MCacheSys:     (rtm.MCacheSys),
//		MSpanInuse:    (rtm.MSpanInuse),
//		MSpanSys:      (rtm.MSpanSys),
//		Mallocs:       (rtm.Mallocs),
//		NextGC:        (rtm.NextGC),
//		NumForcedGC:   (rtm.NumForcedGC),
//		NumGC:         (rtm.NumGC),
//		OtherSys:      (rtm.OtherSys),
//		PauseTotalNs:  (rtm.PauseTotalNs),
//		StackInuse:    (rtm.StackInuse),
//		StackSys:      (rtm.StackSys),
//		Sys:           (rtm.Sys),
//		TotalAlloc:    (rtm.TotalAlloc),
//		RandomValue:   rand.(),
//		PollCount:     1,
//	}
//}

//func (m *Metrics) Update() {
//	var rtm runtime.MemStats
//
//	for {
//		<-time.After(pollInterval)
//		runtime.ReadMemStats(&rtm)
//
//		m.Alloc = (rtm.Alloc)
//		m.BuckHashSys = (rtm.BuckHashSys)
//		m.Frees = rtm.Frees
//
//
//			Frees:         (rtm.Frees),
//			GCCPUFraction: rtm.GCCPUFraction,
//			GCSys:         (rtm.GCSys),
//			HeapAlloc:     (rtm.HeapAlloc),
//			HeapIdle:      (rtm.HeapIdle),
//			HeapInuse:     (rtm.HeapInuse),
//			HeapObjects:   (rtm.HeapObjects),
//			HeapReleased:  (rtm.HeapReleased),
//			HeapSys:       (rtm.HeapSys),
//			LastGC:        (rtm.LastGC),
//			Lookups:       (rtm.Lookups),
//			MCacheInuse:   (rtm.MCacheInuse),
//			MCacheSys:     (rtm.MCacheSys),
//			MSpanInuse:    (rtm.MSpanInuse),
//			MSpanSys:      (rtm.MSpanSys),
//			Mallocs:       (rtm.Mallocs),
//			NextGC:        (rtm.NextGC),
//			NumForcedGC:   (rtm.NumForcedGC),
//			NumGC:         (rtm.NumGC),
//			OtherSys:      (rtm.OtherSys),
//			PauseTotalNs:  (rtm.PauseTotalNs),
//			StackInuse:    (rtm.StackInuse),
//			StackSys:      (rtm.StackSys),
//			Sys:           (rtm.Sys),
//			TotalAlloc:    (rtm.TotalAlloc),
//			RandomValue:   rand.(),
//			PollCount:     1,
//	}
//}
