package main

import "testing"

func Test_subarraySum(t *testing.T) {
	type args struct {
		arr []int
		sum int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"`sum` number exists",
			args{[]int{1, 2, 3}, 5},
			true,
		},
		{
			"`sum` number exists",
			args{[]int{3, 2, 5, 1}, 6},
			true,
		},
		{
			"`sum` number does not exists",
			args{[]int{11, 21, 4}, 5},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := subarraySum(tt.args.arr, tt.args.sum); got != tt.want {
				t.Errorf("subarraySum() = %v, want %v", got, tt.want)
			}
		})
	}
}
