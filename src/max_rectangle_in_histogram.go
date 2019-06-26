package main

// [3, 1, 4, 4, 3, 2, 2, 2, 2]
// result: 6
// 2, 3, 2, 0, 1
// solution: 

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func maxRectInHist(nums []int) int {
	history := []int{}
	
	for i := 0; i+1 < len(nums) ; i++ {
		minValue := min(nums[i], nums[i+1])
		if len(history) == 0 || history[len(history)-1] != minValue {
			history = append(history, minValue)
			history = append(history, minValue)
		} else if history[len(history)-1] == minValue {
			history = append(history, minValue)
		}
	}

	valuesMap := map[int]int{}
	for i := 0; i < len(history); i++ {
		valuesMap[history[i]] += history[i] 
	}

	result := 0 
	for _, v := range valuesMap {
		if result < v {
			result = v
		}
	}

	return result
}

