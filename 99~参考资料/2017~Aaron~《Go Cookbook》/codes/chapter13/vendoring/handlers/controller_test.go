package handlers

import (
	"reflect"
	"testing"

	"github.com/agtorre/go-cookbook/chapter13/vendoring/models"
)

func TestNewController(t *testing.T) {
	type args struct {
		db models.DB
	}
	tests := []struct {
		name string
		args args
		want *Controller
	}{
		{"base-case", args{models.NewDB()}, &Controller{models.NewDB()}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewController(tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewController() = %v, want %v", got, tt.want)
			}
		})
	}
}

type mockdb struct {
	mockGetScore func() (int64, error)
	mockSetScore func(int64) error
}

// GetScore returns the score atomically
func (d *mockdb) GetScore() (int64, error) {
	if d.mockGetScore != nil {
		return d.mockGetScore()
	}
	return 0, nil
}

// SetScore stores a new value atomically
func (d *mockdb) SetScore(score int64) error {
	if d.mockSetScore != nil {
		return d.mockSetScore(score)
	}
	return nil
}
