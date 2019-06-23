package main

import (
	"reflect"
	"testing"
)

func Test_nextLargerNumber(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"NextLargerNumber#1",
			args{[]int{3,1,3,4}},
			[]int{4,3,4,-1},
		},
		{
			"NextLargerNumber#2 (Edge case)",
			args{[]int{}},
			[]int{},
		},
		{
			"NextLargerNumber#3",
			args{[]int{1,2,3,4}},
			[]int{2,3,4,-1},
		},
		{
			"NextLargerNumber#4",
			args{[]int{3,5,3,4}},
			[]int{5,-1,4,5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nextLargerNumber(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("nextLargerNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
