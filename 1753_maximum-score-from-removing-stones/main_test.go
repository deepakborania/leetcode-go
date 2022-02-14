package main

import "testing"

func Test_maximumScore(t *testing.T) {
	type args struct {
		a int
		b int
		c int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"Test 1",
			args{a: 2, b: 4, c: 6},
			6,
		},
		{
			"Test 2",
			args{a: 4, b: 4, c: 6},
			7,
		},
		{
			"Test 3",
			args{a: 1, b: 8, c: 8},
			8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumScore(tt.args.a, tt.args.b, tt.args.c); got != tt.want {
				t.Errorf("maximumScore() = %v, want %v", got, tt.want)
			}
		})
	}
}
