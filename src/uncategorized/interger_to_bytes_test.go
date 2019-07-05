package main

import (
	"reflect"
	"testing"
)

func Test_decimalToBinary(t *testing.T) {
	type args struct {
		val uint64
	}
	tests := []struct {
		name string
		args args
		want []uint64
	}{
		{
			"decimalToBinary#1",
			args{3},
			[]uint64{1,1},
		},
		{
			"decimalToBinary#2",
			args{15},
			[]uint64{1,1,1,1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := decimalToBinary(tt.args.val); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("decimalToBinary() = %v, want %v", got, tt.want)
			}
		})
	}
}
