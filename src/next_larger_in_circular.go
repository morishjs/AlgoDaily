package main

// A circular array is one where the next element of the last element is the first element:

// Instead of [1, 2, 3, 4], imagine the following, where after index 7, we'd move back to index 0.

// Can you write a method nextLargerNumber(nums: array) to print the next immediate larger number for every element in the array? Note: for any element within the circular array, the next immediate larger number could be found circularly-- past the end and before it. If there is no number greater, return -1.

// Take the following analysis for each index:

// nextLargerNumber([3, 1, 3, 4])

// // [4, 3, 4, -1]

// // 3's next greater is 4

// // 1's next greater is 3

// // 3's next greater is 4 again

// // 4 will look to the beginning, but find nothing, so -1

// solution: https://algodaily.com/challenges/next-larger-in-a-circular-array

// args{[]int{3,1,3,4}},
// []int{4,3,4,-1},

func nextLargerNumber(arr []int) []int {
	if len(arr) == 0 {
		return []int{}
	}

	result := []int{}
	stack := []int{}

	for i := 0; i < len(arr); i++ {
		result = append(result, -1)
	}

	for i := 0; i < len(arr)*2; i++ {
		currentValue := arr[i%len(arr)]

		for len(stack) != 0 && arr[stack[len(stack)-1]] < currentValue {
			result[stack[len(stack)-1]] = currentValue
			stack = stack[:len(stack)-1]
		}

		if i < len(arr) {
			stack = append(stack, i)
		}
	}

	return result
}
