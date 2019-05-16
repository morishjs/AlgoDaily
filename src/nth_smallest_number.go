package main

import (
	"fmt"
	ds "go-algo-daily/src/data_structures"
	"math/rand"
	"time"
)

func randomNumberGenerator(upper int) func() int {
	rand.Seed(time.Now().UnixNano())

	return func() int {
		return rand.Intn(upper)
	}
}

// 10 2 9 + 7
// 7 2 9
// 9 2 7 (heapify)

// 9 2 7 + 8
// 8 2 7 (heapify)

// 8 2 7 + 9
// 8 2 7
func main() {
	nthNum := 3
	randNumGenerator := randomNumberGenerator(10)
	
	maxHeap := ds.MaxHeap{}

	var nthArr []int
	count := 0

	for index := 0; index < 15; index++ {
		n := randNumGenerator()
		fmt.Println("Generated number is ", n)

		if count < nthNum {
			maxHeap.Add(n)
			count++
		} else if count == nthNum {
			nthArr = maxHeap.Nodes
			count++
		} else {
			prevNthSmallestNumber := maxHeap.GetRootValue()
			
			if prevNthSmallestNumber > n {
				nthArr[0] = n
				heap := ds.NewMaxHeap(nthArr)

				nthArr = heap.Nodes
			}

			fmt.Printf("%dth smallest number is %d\n", nthNum, nthArr[0])
		}
	}
}
