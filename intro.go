package main

import "fmt"

func main() {
	fmt.Print("Running program...")
}

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

func summed(nums []int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}
