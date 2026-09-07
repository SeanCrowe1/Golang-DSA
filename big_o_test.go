package main

import (
	"fmt"
	"testing"
)

func TestFindMax(t *testing.T) {
	type testCase struct {
		nums     []float64
		expected float64
	}
	testCases := []testCase{
		{[]float64{7, 4, 3, 100, 2343243, 343434, 1, 2, 32}, 2343243},
		{[]float64{12, 12, 12}, 12},
		{[]float64{10, 200, 3000, 5000, 4}, 5000},
		{[]float64{0}, 0},
		{[]float64{-1, -2, -3}, -1},
		{[]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 10},
		{[]float64{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}, 10},
	}

	passCount := 0
	failCount := 0

	for _, test := range testCases {
		output := findMax(test.nums)
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

func TestDoesNameExist(t *testing.T) {
	type testCase struct {
		firstNames int
		lastNames  int
		fullName   string
		expected   bool
	}
	testCases := []testCase{
		{100, 100, "bob0 gonzalez0", true},
		{500, 500, "maria1 smith1", true},
		{1000, 1000, "bob500 smith1", false},
		{2000, 2000, "bob1999 wagner1998", false},
		{3000, 3000, "sally2999 smith2998", true},
	}

	passCount := 0
	failCount := 0

	for _, test := range testCases {
		firstNames := getFirstNames(test.firstNames)
		lastNames := getLastNames(test.lastNames)
		output := doesNameExist(firstNames, lastNames, test.fullName)
		if fmt.Sprintf("%t", output) != fmt.Sprintf("%t", test.expected) {
			failCount++
			t.Errorf(`---------------------------------
Inputs:     (%v, %v, %v)
Expecting:  %t
Actual:     %t
Fail
`, test.firstNames, test.lastNames, test.fullName, test.expected, output)
		} else {
			passCount++
			fmt.Printf(`---------------------------------
Inputs:     (%v, %v, %v)
Expecting:  %t
Actual:     %t
Pass
`, test.firstNames, test.lastNames, test.fullName, test.expected, output)
		}
	}

	fmt.Println("---------------------------------")
	fmt.Printf("%d passed, %d failed\n", passCount, failCount)
}

func getFirstNames(num int) []string {
	names := []string{}
	for i := 0; i < num; i++ {
		m := i % 3
		switch m {
		case 0:
			names = append(names, fmt.Sprintf("bob%d", i))
		case 1:
			names = append(names, fmt.Sprintf("maria%d", i))
		case 2:
			names = append(names, fmt.Sprintf("sally%d", i))
		}
	}
	return names
}

func getLastNames(num int) []string {
	names := []string{}
	for i := 0; i < num; i++ {
		m := i % 3
		switch m {
		case 0:
			names = append(names, fmt.Sprintf("gonzalez%d", i))
		case 1:
			names = append(names, fmt.Sprintf("smith%d", i))
		case 2:
			names = append(names, fmt.Sprintf("wagner%d", i))
		}
	}
	return names
}
