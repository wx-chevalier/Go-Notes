package handlers

import (
	"errors"
	"net/http"
	"testing"

	"net/http/httptest"

	"github.com/agtorre/go-cookbook/chapter13/vendoring/models"
)

func TestController_GetHandler(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		c    *Controller
		args args
	}{
		{"base-case", NewController(models.NewDB()), args{httptest.NewRecorder(), httptest.NewRequest("GET", "/", nil)}},
		{"get error", NewController(&mockdb{
			mockGetScore: func() (int64, error) {
				return 0, errors.New("failed")
			},
		}), args{httptest.NewRecorder(), httptest.NewRequest("GET", "/", nil)}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.c.GetHandler(tt.args.w, tt.args.r)
		})
	}
}
