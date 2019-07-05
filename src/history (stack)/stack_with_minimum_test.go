package main

import (
	"testing"
)

func TestMinStack_push(t *testing.T) {
	type fields struct {
		minStack []int
		minValues []int
	}
	type args struct {
		val int
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		minValues int
	}{
		{
			"Push a value",
			fields{[]int{2, 3, 4}, []int{2, 2, 2}},
			args{1},
			1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &MinStack{
				minStack: tt.fields.minStack,
				minValues: tt.fields.minValues,
			}
			s.push(tt.args.val)

			if s.minValues[len(s.minValues)-1] != tt.minValues {
				t.Errorf("MinStack.push(), actual: %v, expected: %v", s.minValues, tt.minValues)
				return
			}
		})
	}
}

func TestMinStack_pop(t *testing.T) {
	type fields struct {
		minStack []int
		minValues []int
	}
	tests := []struct {
		name     string
		fields   fields
		want     int
		minValues int
	}{
		{
			"Stack Pop",
			fields{[]int{2, 3, 4}, []int{2, 2, 2}},
			4,
			2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &MinStack{
				minStack: tt.fields.minStack,
				minValues: tt.fields.minValues,
			}
			got, err := s.pop()
			if err != nil {
				t.Errorf("MinStack.pop() error = %v", err)
				return
			}
			if got != tt.want {
				t.Errorf("MinStack.pop() = %v, want %v", got, tt.want)
			}
			if tt.minValues != s.minValues[len(s.minValues)-1] {
				t.Error("MinStack.pop() unexpected minimum value in stack")
			}
		})
	}
}

func TestMinStack_peek(t *testing.T) {
	type fields struct {
		minStack []int
		minValues []int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			"peek()",
			fields{[]int{2, 3, 4}, []int{2, 2, 2}},
			4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &MinStack{
				minStack: tt.fields.minStack,
				minValues: tt.fields.minValues,
			}
			got, _ := s.peek()
			if got != tt.want {
				t.Errorf("MinStack.peek() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMinStack_empty(t *testing.T) {
	type fields struct {
		minStack []int
		minValues []int
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			"empty()",
			fields{[]int{2, 3, 4}, []int{2, 2, 2}},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &MinStack{
				minStack: tt.fields.minStack,
				minValues: tt.fields.minValues,
			}
			if got := s.empty(); got != tt.want {
				t.Errorf("MinStack.empty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMinStack_min(t *testing.T) {
	type fields struct {
		minStack []int
		minValues []int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			"min()",
			fields{[]int{2, 3, 4}, []int{2, 2, 2}},
			2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &MinStack{
				minStack: tt.fields.minStack,
				minValues: tt.fields.minValues,
			}
			got, _ := s.min()
			if got != tt.want {
				t.Errorf("MinStack.min() = %v, want %v", got, tt.want)
			}
		})
	}
}
