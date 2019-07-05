// Given an array of numbers, return true if there is a contiguous subarray that sums up to a certain number n.

// const arr = [1, 2, 3], sum = 5
// subarraySum(arr, sum)
// // true

// const arr = [11, 21, 4], sum = 9
// subarraySum(arr, sum)
// // false
// In the above examples, 2, 3 sum up to 5 so we return true. On the other hand, no subarray in [11, 21, 4] can sum up to 9.

// solution: https://algodaily.com/challenges/contiguous-subarray-sum

package main

func subarraySum(arr []int, sum int) bool {
	for i, v := range arr {
		if i+1 < len(arr) && v+arr[i+1] == sum {
			return true
		}
	}

	return false
}
