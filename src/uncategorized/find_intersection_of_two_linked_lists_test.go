package main

import (
	ds "go-algo-daily/src/data_structures"
	"reflect"
	"testing"
)

func initList(values []int) *ds.List {
	l := new(ds.List)
	for _, v := range values {
		l.Append(v)
	}

	return l
}

func Test_getIntersection(t *testing.T) {
	type args struct {
		listOne *ds.List
		listTwo *ds.List
	}
	tests := []struct {
		name string
		args args
		want *ds.List
	}{
		{
			"GetIntersection#1",
			args{initList([]int{1,2,3,4,5}), initList([]int{1,2,3,4,5})},
			initList([]int{1,2,3,4,5}),
		},
		{
			"GetIntersection#2",
			args{initList([]int{1,2,3,5,6}), initList([]int{1,2,3,4,5})},
			initList([]int{1,2,3,5}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getIntersection(tt.args.listOne, tt.args.listTwo); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getIntersection() = %v, want %v", got, tt.want)
			}
		})
	}
}
