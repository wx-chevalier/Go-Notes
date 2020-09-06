package main

import (
	"reflect"
	"testing"
)

func Test_FlattenArray(t *testing.T) {
	type args struct {
		l []interface{}
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			"name",
			args{[]interface{}{2, 1, []interface{}{3, []interface{}{4, 5}, 6}, 7, []interface{}{8}}},
			[]int{2, 1, 3, 4, 5, 6, 7, 8},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FlattenArray(tt.args.l); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FlattenArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
