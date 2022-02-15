package main

import (
	"reflect"
	"testing"
)

func Test_findClosestElements(t *testing.T) {
	type args struct {
		arr []int
		k   int
		x   int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"Test 1", args{[]int{1, 2, 3, 4, 5}, 4, 3}, []int{1, 2, 3, 4}},
		{"Test 2", args{[]int{1, 2, 3, 4, 5}, 4, -1}, []int{1, 2, 3, 4}},
		{"Test 3", args{[]int{1, 3, 3, 3, 3, 3, 3, 10}, 3, 3}, []int{3, 3, 3}},
		{"Test 4", args{[]int{1, 2}, 1, 1}, []int{1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findClosestElements(tt.args.arr, tt.args.k, tt.args.x); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findClosestElements() = %v, want %v", got, tt.want)
			}
		})
	}
}
