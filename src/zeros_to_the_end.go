// Write a method that moves all zeros in an array to its end. You should maintain the order of all other elements. Here's an example:

// Example #1
// zerosToEnd([1, 0, 0, 0, 4, 2])

// Exaple #2
// zerosToEnd([1, 0, 2, 0, 4, 0])

// [1, 2, 0, 0, 4, 0]

// [1, 2, 0, 0, 4, 0]

// // [1, 2, 4, 0, 0, 0]

// And another one:

// zerosToEnd([1, 0, 2, 0, 4, 0])

// // [1, 2, 4, 0, 0, 0]

// Fill in the following function signature:

// function zerosToEnd(nums) {

//   return;

// };

// Can you do this without instantiating a new array?

// solution: https://algodaily.com/challenges/zeros-to-the-end

package main

func zerosToEnd(nums []int) []int {
	var changeIndex int

	for i,v := range nums {
		if v == 0 {
			changeIndex = i
			break
		}
	}

	for currentIndex, v := range nums {
		if v != 0 && changeIndex < currentIndex {
			nums[changeIndex] = v
			nums[currentIndex] = 0

			changeIndex += 1
		}
	}

	return nums
}
