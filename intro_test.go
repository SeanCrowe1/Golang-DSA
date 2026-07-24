package main

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	type testCase struct {
		nums     []int
		expected int
	}
	runCases := []testCase{
		{[]int{7, 4, 3, 100, 2343243, 343434, 1, 2, 32}, 1},
		{[]int{12, 12, 12}, 12},
		{[]int{10, 200, 3000, 5000, 4}, 4},
		{[]int{1}, 1},
		{[]int{1, 2, 3, 4, 5}, 1},
		{[]int{5, 4, 3, 2, 1}, 1},
		{[]int{100, 200, 300, 400, 500}, 100},
		{[]int{500, 400, 300, 200, 100}, 100},
		{[]int{}, 0},
	}

	passCount := 0
	failCount := 0

	for _, test := range runCases {
		output := findMinimum(test.nums)
		if fmt.Sprintf("%d", output) != fmt.Sprintf("%d", test.expected) {
			failCount++
			t.Errorf(`---------------------------------
Inputs:     (%v)
Expecting:  %d
Actual:     %d
Fail
`, test.nums, test.expected, output)
		} else {
			passCount++
			fmt.Printf(`---------------------------------
Inputs:     (%v)
Expecting:  %d
Actual:     %d
Pass
`, test.nums, test.expected, output)
		}
	}

	fmt.Println("---------------------------------")
	fmt.Printf("%d passed, %d failed\n", passCount, failCount)
}
