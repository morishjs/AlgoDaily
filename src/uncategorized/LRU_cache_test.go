package main

import (
	"testing"
)

func TestCache_Put(t *testing.T) {
	type args struct {
		index int
		value int
	}
	tests := []struct {
		name string
		size int
		args []args
	}{
		{
			"Put",
			3,
			[]args{args{1, 1}, args{2, 4}, args{3, 9}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cache := NewCache(tt.size)
			for _, arg := range tt.args {
				cache.Put(arg.index, arg.value)
			}
		})
	}
}

func TestCache_Get(t *testing.T) {
	type args struct {
		index int
		value int
	}
	tests := []struct {
		name string
		size int
		args []args
		getIndex int
		emptyIndex int
		want int
	}{
		{
			"Get",
			3,
			[]args{args{1, 1}, args{2, 4}, args{3, 9}},
			1,
			2,
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cache := NewCache(tt.size)
			for _, arg := range tt.args {
				cache.Put(arg.index, arg.value)
			}

			cache.Get(tt.getIndex)
			cache.Put(4, 16)

			if cache.Get(tt.emptyIndex) != tt.want {
				t.Errorf("Get() = %v, want %v", cache.Get(tt.emptyIndex), tt.want)
			}
		})
	}
}
