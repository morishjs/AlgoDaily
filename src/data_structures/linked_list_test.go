package ds

import (
	"reflect"
	"testing"
)

func TestList_Prepend(t *testing.T) {
	type fields struct {
		list *List
	}
	type args struct {
		values []int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []int
	}{
		{
			"Prepend",
			fields{new(List)},
			args{[]int{1, 2, 3}},
			[]int{3, 2, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := tt.fields.list
			for _, v := range tt.args.values {
				list.Prepend(v)
			}

			if got := list.ToArray(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("List.Prepend() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestList_Append(t *testing.T) {
	type fields struct {
		list *List
	}
	type args struct {
		values []int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []int
	}{
		{
			"Append",
			fields{new(List)},
			args{[]int{1, 2, 3}},
			[]int{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := tt.fields.list
			for _, v := range tt.args.values {
				list.Append(v)
			}

			if got := list.ToArray(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("List.Prepend() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestList_Contains(t *testing.T) {
	type fields struct {
		list *List
	}
	type args struct {
		values []int
		value int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			"Contains",
			fields{new(List)},
			args{[]int{1, 2, 3}, 1},
			true,
		},
		{
			"Do not contain",
			fields{new(List)},
			args{[]int{1, 2, 3}, 4},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := tt.fields.list
			for _, v := range tt.args.values {
				list.Append(v)
			}
			if got := list.Contains(tt.args.value); got != tt.want {
				t.Errorf("List.Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestList_lazyInit(t *testing.T) {
	type fields struct {
		root *ListNode
		len  int
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := &List{
				root: tt.fields.root,
				len:  tt.fields.len,
			}
			list.lazyInit()
		})
	}
}

func TestList_Init(t *testing.T) {
	type fields struct {
		root *ListNode
		len  int
	}
	tests := []struct {
		name   string
		fields fields
		want   *List
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := &List{
				root: tt.fields.root,
				len:  tt.fields.len,
			}
			if got := list.Init(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("List.Init() = %v, want %v", got, tt.want)
			}
		})
	}
}
