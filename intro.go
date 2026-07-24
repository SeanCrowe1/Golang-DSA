package main

func findMinimum(nums []int) int {
	if len(nums) == 0 {
		return 0
	} else if len(nums) == 1 {
		return nums[0]
	}

	min := nums[0]
	for _, num := range nums[1:] {
		if num < min {
			min = num
		}
	}

	return min
}
