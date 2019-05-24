// Challenge: Nth Smallest Number in a Stream

// Many modern programming languages have the notion of a stream. A stream is a resource that's broken down and sent (usually over a network) in small chunks.

// Imagine a video that's playing-- it would be inefficient to need to download an entire 1 gigabyte movie before being able to watch it. Instead, browsers will gradually download it in pieces.

// Given a stream of numbers with an infinite length, can you find the nth smallest element at each point of the streaming? In the below example, we are looking for the second smallest element.

// // stream is the resource, the "..." represents future integers

// const stream = [3, 4, 5, 6, 7, ...]

// nthSmallestNumber(2); // 2nd smallest number

// // [null, 4, 4, 4, ...] (4 is always second smallest)

// Let's use a for-loop to simulate a stream. Your method will be called as follows:

// for (let i = 0; i < 10; i++) {

//   const newInt = Math.floor(Math.random() * 10) + 1;

//   nthSmallestNumber(n);

//   // return nth smallest number

// }

// Solution: https://algodaily.com/challenges/nth-smallest-number-in-a-stream

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
