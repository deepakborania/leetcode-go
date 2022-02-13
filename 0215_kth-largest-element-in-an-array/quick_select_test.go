package main

import "testing"

func Test_findKthLargest2(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test 1",
			args: args{nums: []int{3, 2, 1, 5, 6, 4}, k: 2},
			want: 5,
		},
		{
			name: "Test 2",
			args: args{nums: []int{3, 2, 3, 1, 2, 4, 5, 5, 6}, k: 4},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findKthLargest2(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("findKthLargest2() = %v, want %v", got, tt.want)
			}
		})
	}
}
