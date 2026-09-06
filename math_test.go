package main

import (
	"fmt"
	"math"
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

func TestGetFollowerPrediction(t *testing.T) {
	type testCase struct {
		followerCount  int
		influencerType string
		numMonths      int
		expected       int
	}
	testCases := []testCase{
		{10, "fitness", 1, 40},
		{10, "fitness", 2, 160},
		{12, "cosmetic", 4, 972},
		{15, "business", 4, 240},
		{10, "fitness", 5, 10240},
		{10, "fitness", 6, 40960},
		{10, "fitness", 7, 163840},
		{10, "fitness", 8, 655360},
		{10, "tech", 9, 5120},
	}

	passCount := 0
	failCount := 0

	for _, test := range testCases {
		output := getFollowerPrediction(test.followerCount, test.influencerType, test.numMonths)
		if fmt.Sprintf("%d", output) != fmt.Sprintf("%d", test.expected) {
			failCount++
			t.Errorf(`---------------------------------
Inputs:     (%v, %v, %v)
Expecting:  %d
Actual:     %d
Fail
`, test.followerCount, test.influencerType, test.numMonths, test.expected, output)
		} else {
			passCount++
			fmt.Printf(`---------------------------------
Inputs:     (%v, %v, %v)
Expecting:  %d
Actual:     %d
Pass
`, test.followerCount, test.influencerType, test.numMonths, test.expected, output)
		}
	}

	fmt.Println("---------------------------------")
	fmt.Printf("%d passed, %d failed\n", passCount, failCount)
}

func TestGetInfluencerScore(t *testing.T) {
	type testCase struct {
		followerCount               int
		averageEngagementPercentage float64
		expected                    float64
	}
	testCases := []testCase{
		{40000, 0.3, 5.0},
		{43000, 0.1, 2.0},
		{100000, 0.6, 10.0},
		{1, 1, 0.0},
		{200, 0.8, 6.0},
		{300000, 0.5, 9.0},
		{500000, 0.2, 4.0},
		{750000, 0.7, 14.0},
	}

	passCount := 0
	failCount := 0

	for _, test := range testCases {
		output := math.RoundToEven(getInfluencerScore(test.followerCount, test.averageEngagementPercentage))
		if fmt.Sprintf("%.2f", output) != fmt.Sprintf("%.2f", test.expected) {
			failCount++
			t.Errorf(`---------------------------------
Inputs:     (%v, %v)
Expecting:  %.2f
Actual:     %.2f
Fail
`, test.followerCount, test.averageEngagementPercentage, test.expected, output)
		} else {
			passCount++
			fmt.Printf(`---------------------------------
Inputs:     (%v, %v)
Expecting:  %.2f
Actual:     %.2f
Pass
`, test.followerCount, test.averageEngagementPercentage, test.expected, output)
		}
	}

	fmt.Println("---------------------------------")
	fmt.Printf("%d passed, %d failed\n", passCount, failCount)
}

func TestNumPossibleOrders(t *testing.T) {
	type testCase struct {
		numPosts int
		expected int
	}
	testCases := []testCase{
		{2, 2},
		{3, 6},
		{5, 120},
		{1, 1},
		{6, 720},
		{7, 5040},
		{8, 40320},
		{9, 362880},
		{11, 39916800},
	}

	passCount := 0
	failCount := 0

	for _, test := range testCases {
		output := numPossibleOrders(test.numPosts)
		if fmt.Sprintf("%d", output) != fmt.Sprintf("%d", test.expected) {
			failCount++
			t.Errorf(`---------------------------------
Inputs:     (%v)
Expecting:  %d
Actual:     %d
Fail
`, test.numPosts, test.expected, output)
		} else {
			passCount++
			fmt.Printf(`---------------------------------
Inputs:     (%v)
Expecting:  %d
Actual:     %d
Pass
`, test.numPosts, test.expected, output)
		}
	}

	fmt.Println("---------------------------------")
	fmt.Printf("%d passed, %d failed\n", passCount, failCount)
}

func TestDecayedFollowers(t *testing.T) {
	type testCase struct {
		initialFollowers  int
		fractionLostDaily float64
		days              int
		expected          float64
	}
	testCases := []testCase{
		{200, 0.5, 1, 100},
		{200, 0.4, 2, 72},
		{200, 0.05, 3, 171},
		{1000, 0.005, 2, 990},
		{1000, 0.05, 3, 857},
		{1200, 0.55, 8, 2},
		{1200, 0.09, 16, 265},
		{0, 0.5, 1, 0},
		{100, 0, 5, 100},
	}

	passCount := 0
	failCount := 0

	for _, test := range testCases {
		output := math.RoundToEven(decayedFollowers(test.initialFollowers, test.fractionLostDaily, test.days))
		if fmt.Sprintf("%.2f", output) != fmt.Sprintf("%.2f", test.expected) {
			failCount++
			t.Errorf(`---------------------------------
Inputs:     (%v, %v, %v)
Expecting:  %.2f
Actual:     %.2f
Fail
`, test.initialFollowers, test.fractionLostDaily, test.days, test.expected, output)
		} else {
			passCount++
			fmt.Printf(`---------------------------------
Inputs:     (%v, %v, %v)
Expecting:  %.2f
Actual:     %.2f
Pass
`, test.initialFollowers, test.fractionLostDaily, test.days, test.expected, output)
		}
	}

	fmt.Println("---------------------------------")
	fmt.Printf("%d passed, %d failed\n", passCount, failCount)
}

func TestLogScale(t *testing.T) {
	type testCase struct {
		data     []float64
		base     float64
		expected []float64
	}
	testCases := []testCase{
		{[]float64{1, 10, 100, 1000}, 10, []float64{0.0, 1.0, 2.0, 3.0}},
		{[]float64{1, 2, 4, 8}, 2, []float64{0.0, 1.0, 2.0, 3.0}},
		{[]float64{2, 4, 8, 16}, 2, []float64{1.0, 2.0, 3.0, 4.0}},
		{[]float64{3, 9, 27, 81}, 3, []float64{1.0, 2.0, 3.0, 4.0}},
		{[]float64{5, 25, 125, 625}, 5, []float64{1.0, 2.0, 3.0, 4.0}},
		{[]float64{10, 100, 1000, 10000}, 10, []float64{1.0, 2.0, 3.0, 4.0}},
		{[]float64{20, 400, 8000, 160000}, 20, []float64{1.0, 2.0, 3.0, 4.0}},
	}

	passCount := 0
	failCount := 0

	for _, test := range testCases {
		output := logScale(test.data, test.base)
		if fmt.Sprintf("%.2f", output) != fmt.Sprintf("%.2f", test.expected) {
			failCount++
			t.Errorf(`---------------------------------
Inputs:     (%v, %v)
Expecting:  %v
Actual:     %v
Fail
`, test.data, test.base, test.expected, output)
		} else {
			passCount++
			fmt.Printf(`---------------------------------
Inputs:     (%v, %v)
Expecting:  %v
Actual:     %v
Pass
`, test.data, test.base, test.expected, output)
		}
	}

	fmt.Println("---------------------------------")
	fmt.Printf("%d passed, %d failed\n", passCount, failCount)
}

func TestAverageFollowers(t *testing.T) {
	type testCase struct {
		nums     []int
		expected float64
	}
	testCases := []testCase{
		{[]int{1}, 1},
		{[]int{1, 2, 3, 4, 5, 6, 7}, 4},
		{[]int{12, 12, 12}, 12},
		{[]int{}, 0},
		{[]int{0}, 0},
		{[]int{100, 200, 300, 400, 500}, 300},
		{[]int{5, 10, 200, 3000, 5000}, 1643},
		{[]int{12_345, 618_222, 58_832_221, 2_180_831_475, 8_663_863_102}, 2180831473},
	}

	passCount := 0
	failCount := 0

	for _, test := range testCases {
		output := math.RoundToEven(averageFollowers(test.nums))
		if fmt.Sprintf("%.2f", output) != fmt.Sprintf("%.2f", test.expected) {
			failCount++
			t.Errorf(`---------------------------------
Inputs:     (%v)
Expecting:  %v
Actual:     %v
Fail
`, test.nums, test.expected, output)
		} else {
			passCount++
			fmt.Printf(`---------------------------------
Inputs:     (%v)
Expecting:  %v
Actual:     %v
Pass
`, test.nums, test.expected, output)
		}
	}

	fmt.Println("---------------------------------")
	fmt.Printf("%d passed, %d failed\n", passCount, failCount)
}
