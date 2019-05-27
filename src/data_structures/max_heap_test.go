package ds

import "testing"

func TestMaxHeap_Add(t *testing.T) {
	type fields struct {
		nodes []int
	}
	tests := []struct {
		name   string
		fields fields
		values []int
	}{
		{
			"Add",
			fields{[]int{1}},
			[]int{4, 3, 52, 2, 5, 13},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			heap := &MaxHeap{
				Nodes: tt.fields.nodes,
			}
			for _, v := range tt.values {
				heap.Add(v)	
			}
			
			heap.Print()
		})
	}
}
