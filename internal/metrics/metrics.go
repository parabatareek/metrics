package metrics

import (
	"reflect"
	"runtime"
	"time"
)

const (
	pollInterval   = 2 * time.Second
	reportInterval = 10 * time.Second
)

func GetMetricsName() []string {
	return []string{"Alloc", "BuckHashSys", "Frees", "GCCPUFraction", "GCSys", "HeapAlloc", "HeapIdle", "HeapInuse", "HeapObjects", "HeapReleased", "HeapSys", "LastGC", "Lookups", "MCacheInuse", "MCacheSys", "MSpanInuse", "MSpanSys", "Mallocs", "NextGC", "NumForcedGC", "NumGC", "OtherSys", "PauseTotalNs", "StackInuse", "StackSys", "Sys", "TotalAlloc"}
}

func GetMetrics() (metrics map[string]float64) {
	rtm := runtime.MemStats{}
	runtime.ReadMemStats(&rtm)
	statsVal := reflect.ValueOf(rtm)
	statsType := reflect.TypeOf(rtm)

	for _, val := range GetMetricsName() {
		for i := 0; i < statsVal.NumField(); i++ {
			if val == statsType.Field(i).Name {
				metrics[val] = statsVal.Field(i).Float()
			}
		}
	}
	return metrics
}

//
//type Metrics struct {
//	Alloc
//	BuckHashSys
//	Frees
//	GCCPUFraction
//	GCSys
//	HeapAlloc
//	HeapIdle
//	HeapInuse
//	HeapObjects
//	HeapReleased
//	HeapSys
//	LastGC
//	Lookups
//	MCacheInuse
//	MCacheSys
//	MSpanInuse
//	MSpanSys
//	Mallocs
//	NextGC
//	NumForcedGC
//	NumGC
//	OtherSys
//	PauseTotalNs
//	StackInuse
//	StackSys
//	Sys
//	TotalAlloc
//	RandomValue
//	PollCount     int64
//}
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
