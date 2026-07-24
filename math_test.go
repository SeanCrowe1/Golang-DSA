package main

import (
	"fmt"
	"testing"
)

func TestGetEstimatedSpread(t *testing.T) {
	type testCase struct {
		nums     []int
		expected float64
	}
	testCases := []testCase{
		{[]int{7, 4, 3, 100, 765, 2344, 1, 2, 32}, 5055.91},
		{[]int{12, 12, 12}, 44.85},
		{[]int{10, 200, 3000, 5000, 4}, 11333.1},
		{[]int{}, 0},
		{[]int{1, 1, 1}, 3.74},
		{[]int{100}, 100},
		{[]int{50, 60, 70, 80, 90}, 482.91},
		{[]int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}, 871.69},
		{[]int{5, 10, 15, 20, 25, 30, 35, 40, 45, 50, 55, 60, 65, 70, 75, 80, 85, 90, 95, 100}, 1911.59},
	}

	passCount := 0
	failCount := 0

	for _, test := range testCases {
		output := getEstimatedSpread(test.nums)
		if fmt.Sprintf("%.2f", output) != fmt.Sprintf("%.2f", test.expected) {
			failCount++
			t.Errorf(`---------------------------------
Inputs:     (%v)
Expecting:  %.2f
Actual:     %.2f
Fail
`, test.nums, test.expected, output)
		} else {
			passCount++
			fmt.Printf(`---------------------------------
Inputs:     (%v)
Expecting:  %.2f
Actual:     %.2f
Pass
`, test.nums, test.expected, output)
		}
	}

	fmt.Println("---------------------------------")
	fmt.Printf("%d passed, %d failed\n", passCount, failCount)
}
