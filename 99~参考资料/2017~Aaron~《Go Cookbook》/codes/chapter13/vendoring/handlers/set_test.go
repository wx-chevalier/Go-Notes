package handlers

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/agtorre/go-cookbook/chapter13/vendoring/models"
)

func TestController_SetHandler(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		c    *Controller
		args args
	}{
		{"base-case", NewController(models.NewDB()), args{httptest.NewRecorder(), httptest.NewRequest("GET", "/?score=3", nil)}},
		{"bad-value", NewController(models.NewDB()), args{httptest.NewRecorder(), httptest.NewRequest("GET", "/?score=abc", nil)}},
		{"set error", NewController(&mockdb{
			mockSetScore: func(int64) error {
				return errors.New("failed")
			},
		}), args{httptest.NewRecorder(), httptest.NewRequest("GET", "/?score=1", nil)}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.c.SetHandler(tt.args.w, tt.args.r)
		})
	}
}
