package main

import (
	"reflect"
	"testing"
)

func Test_swap(t *testing.T) {
	type args struct {
		array []int
		x     uint
		y     uint
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			swap(tt.args.array, tt.args.x, tt.args.y)
		})
	}
}

func Test_QuickSort(t *testing.T) {
	type args struct {
		array []int
		start uint
		end   uint
	}

	arr := []int{5, 3, 1, 9, 6, 10, 2}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{"basic", args{arr, 0, uint(len(arr) - 1)}, []int{1, 2, 3, 5, 6, 9, 10}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if QuickSort(tt.args.array, tt.args.start, tt.args.end); !reflect.DeepEqual(arr, tt.want) {
				t.Errorf("QuickSort() = %v, want %v", arr, tt.want)
			}

		})
	}
}
