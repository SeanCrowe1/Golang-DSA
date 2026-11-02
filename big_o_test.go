package main

import (
	"fmt"
	"math/rand"
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

func TestGetAvgBrandFollowers(t *testing.T) {
	type testCase struct {
		numHandles int
		avgAudSize int
		brandName  string
		expected   float64
	}
	testCases := []testCase{
		{10, 1000, "luxa", 437.10},
		{20, 2000, "luxa", 790.05},
		{30, 3000, "luxa", 920.10},
		{40, 4000, "luxa", 1431.58},
		{80, 8000, "luxa", 2685.40},
		{160, 16000, "luxa", 4776.32},
	}

	passCount := 0
	failCount := 0
	rng := rand.New(rand.NewSource(1))

	for _, test := range testCases {
		allHandles := getAllHandles(test.numHandles, test.avgAudSize, rng)
		output := getAvgBrandFollowers(allHandles, test.brandName)
		if fmt.Sprintf("%.2f", output) != fmt.Sprintf("%.2f", test.expected) {
			failCount++
			t.Errorf(`---------------------------------
Inputs:     (%v, %v, %v)
Expecting:  %.2f
Actual:     %.2f
Fail
`, test.numHandles, test.avgAudSize, test.brandName, test.expected, output)
		} else {
			passCount++
			fmt.Printf(`---------------------------------
Inputs:     (%v, %v, %v)
Expecting:  %.2f
Actual:     %.2f
Pass
`, test.numHandles, test.avgAudSize, test.brandName, test.expected, output)
		}
	}

	fmt.Println("---------------------------------")
	fmt.Printf("%d passed, %d failed\n", passCount, failCount)
}

func getAllHandles(num int, audienceSize int, rng *rand.Rand) [][]string {
	allHandles := make([][]string, 0, num)

	minSize := int(float64(audienceSize) - float64(audienceSize)*1.2)
	maxSize := int(float64(audienceSize) + float64(audienceSize)*1.2)

	for range num {
		size := minSize + rng.Intn(maxSize-minSize)
		allHandles = append(allHandles, getUserHandles(size, rng))
	}

	return allHandles
}

func getUserHandles(num int, rng *rand.Rand) []string {
	handles := []string{}
	prefixes := []string{
		"luxaraygirl",
		"theprimerog",
		"luxafanboi",
		"dashlord",
		"saintrex",
		"writergurl",
	}

	for i := range num {
		prefix := prefixes[rng.Intn(len(prefixes))]
		handles = append(handles, fmt.Sprintf("%s%d", prefix, i))
	}

	return handles
}
