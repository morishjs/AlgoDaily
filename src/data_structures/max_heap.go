package ds

import (
	"fmt"
	"math"
)

type MaxHeap struct {
	Nodes []int
}

func NewMaxHeap(arr []int) *MaxHeap{
	heap := MaxHeap{}
	for _, num := range arr {
		heap.Add(num)
	}

	return &heap
}

func (heap *MaxHeap) Add(value int) {
	heap.Nodes = append(heap.Nodes, value)
	heap.Heapify(len(heap.Nodes)-1, value)
}

func (heap *MaxHeap) Heapify(childIndex int, childValue int) {
	parentIndex := getParentIndex(childIndex)

	for childIndex != 0 && heap.Nodes[parentIndex] <= childValue {
		heap.Nodes[parentIndex], heap.Nodes[childIndex] = heap.Nodes[childIndex], heap.Nodes[parentIndex]
		childIndex = parentIndex

		parentIndex = getParentIndex(childIndex)
	}
}

func (heap *MaxHeap) GetRootValue() int {
	return heap.Nodes[0]
}

func (heap *MaxHeap) Print() {
	for i, node := range heap.Nodes {
		fmt.Printf("%d ", node)

		if isEdge(i) {
			fmt.Println()
		}
	}
}

func getParentIndex(currentIndex int) int {
	parentIndex := (currentIndex - 1) / 2

	return int(math.Ceil(float64(parentIndex)))
}

func isEdge(currentIndex int) bool {
	parentIndex := getParentIndex(currentIndex)

	return parentIndex * 2 + 2 == currentIndex
}