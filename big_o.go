package main

import (
	"fmt"
)

// Runtime complexity = O(n)
func findMax(nums []float64) float64 {
	if len(nums) == 0 {
		return 0
	} else if len(nums) == 1 {
		return nums[0]
	}
	maximum := nums[0]
	for _, num := range nums[1:] {
		if num > maximum {
			maximum = num
		}
	}
	return maximum
}

// Runtime complexity = O(n^2)
func doesNameExist(firstNames, lastNames []string, fullName string) bool {
	fmt.Printf("%v %v", firstNames[0], lastNames[0])
	for _, fName := range firstNames {
		for _, lName := range lastNames {
			if fmt.Sprintf("%s %s", fName, lName) == fullName {
				return true
			}
		}
	}
	return false
}
