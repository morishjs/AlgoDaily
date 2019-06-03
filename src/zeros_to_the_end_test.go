package main

import (
	"reflect"
	"testing"
)

func Test_zerosToEnd(t *testing.T) {
	tests := []struct {
		name string
		args []int
		want []int
	}{
		{
			"case#1",
			[]int{1, 0, 2, 0, 4, 0},
			[]int{1, 2, 4, 0, 0, 0},
		},
		{
			"case#2",
			[]int{0, 1, 2, 0, 4, 0},
			[]int{1, 2, 4, 0, 0, 0},
		},
		{
			"case#3",
			[]int{0, 0, 0, 1, 2, 4},
			[]int{1, 2, 4, 0, 0, 0},
		},
		{
			"case#4",
			[]int{1, 2, 4, 0, 0, 0},
			[]int{1, 2, 4, 0, 0, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := zerosToEnd(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("zerosToEnd() = %v, want %v", got, tt.want)
			}
		})
	}
}
