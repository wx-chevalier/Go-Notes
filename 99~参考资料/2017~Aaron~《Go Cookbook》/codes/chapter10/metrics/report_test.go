package metrics

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestReportHandler(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		{"base-case", args{httptest.NewRecorder(), httptest.NewRequest("get", "/", nil)}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ReportHandler(tt.args.w, tt.args.r)
		})
	}
}
