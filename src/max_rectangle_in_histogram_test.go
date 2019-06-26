package main

import "testing"

func Test_maxRectInHist(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"normal case#1",
			args{[]int{3, 1, 4, 4, 3, 2, 2, 2, 2}},
			10,
		},
		{
			"normal case#2",
			args{[]int{3, 1, 4, 2, 2, 1}},
			6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxRectInHist(tt.args.nums); got != tt.want {
				t.Errorf("maxRectInHist() = %v, want %v", got, tt.want)
			}
		})
	}
}
