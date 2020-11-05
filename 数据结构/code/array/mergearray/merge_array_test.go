package mergearray

import (
	"reflect"
	"testing"
)

func TestMergeArray(t *testing.T) {
	type args struct {
		s [][]interface{}
	}

	simpleCase := [][]interface{}{
		[]interface{}{1, 2},
		[]interface{}{3, 4},
	}

	tests := []struct {
		name      string
		args      args
		wantSlice []interface{}
	}{
		{
			"simple",
			args{
				simpleCase,
			},
			[]interface{}{1, 2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotSlice := MergeArray(tt.args.s...); !reflect.DeepEqual(gotSlice, tt.wantSlice) {
				t.Errorf("MergeArray() = %v, want %v", gotSlice, tt.wantSlice)
			}
		})
	}
}
