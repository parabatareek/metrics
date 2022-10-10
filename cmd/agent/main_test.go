package main

import (
	"github.com/parabatareek/metrics.git/internal/metrics"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/url"
	"testing"
)

func Test_getParams(t *testing.T) {
	type args struct {
		dataMetrics *metrics.Metrics
	}
	tests := []struct {
		name string
		args args
		want map[string]string
	}{
		{"Тест типа параметров для request: ",
			args{metrics.NewMetrics()},
			map[string]string{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getParams(tt.args.dataMetrics)
			assert.IsType(t, tt.want, got)
		})
	}
}

func Test_updStats(t *testing.T) {
	type args struct {
		dataMetrics *metrics.Metrics
	}
	tests := []struct {
		name string
		args args
	}{
		{"Тест обновление статиcтики: ",
			args{metrics.NewMetrics()}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var data metrics.Metrics

			assert.NotEqual(t, tt.args.dataMetrics, data)
		})
	}
}

func Test_getRequest(t *testing.T) {
	type want struct {
		method      string
		contentType string
	}
	tests := []struct {
		name string
		want want
	}{
		{"Тест формирования request: ",
			want{http.MethodPost,
				"text/plain"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			urlData := url.Values{}
			urlData.Set("Alloc", "/update/float64/Alloc/100500")
			got := getRequest(&urlData)
			assert.Equal(t, tt.want.method, got.Method)
			assert.Equal(t, tt.want.contentType, got.Header.Get("Content-Type"))
		})
	}
}
