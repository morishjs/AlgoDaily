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
