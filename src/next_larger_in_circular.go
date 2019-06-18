package main

import (
	ds "go-algo-daily/src/data_structures"
	"container/heap"
)

// An IntHeap is a min-heap of ints.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// args{[]int{3,1,3,4}},
// []int{4,3,4,-1},

func nextLargerNumber(arr []int) []int {
	if len(arr) == 0 {
		return []int{}
	}

	numbers := []int{}
	
	list := newCircularList(arr)
	
	elem := list.Root
	for {
		h := &IntHeap{}
		heap.Init(h)

		iter := elem.Next
		
		for elem != iter {
			if elem.Val < iter.Val {
				heap.Push(h, iter.Val)
			}

			iter = iter.Next
		}

		if len(*h) == 0 {
			numbers = append(numbers, -1)
		} else {
			numbers = append(numbers, heap.Pop(h).(int))
		}

		elem = elem.Next
		if list.Root == elem {
			break
		}
	}

	return numbers
}

func newCircularList(arr []int) *ds.List {
	l := new(ds.List)

	for _, v := range arr {
		l.Append(v)
	}

	iter := l.Root
	for iter.Next != nil {
		iter = iter.Next
	}

	iter.Next = l.Root

	return l
}
