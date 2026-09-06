package main

import "math"

func getEstimatedSpread(audiencesFollowers []int) float64 {
	n := float64(len(audiencesFollowers))
	if n == 0 {
		return 0.0
	}

	exp := math.Pow(n, 1.2)

	sum := 0.0
	for _, num := range audiencesFollowers {
		sum += float64(num)
	}

	avg := float64(sum / n)

	estSpread := avg * exp

	return estSpread
}

func getFollowerPrediction(followerCount int, influencerType string, numMonths int) int {
	n := 0
	switch influencerType {
	case "fitness":
		n = 4
	case "cosmetic":
		n = 3
	default:
		n = 2
	}

	return followerCount * int(math.Pow(float64(n), float64(numMonths)))
}

func getInfluencerScore(numFollowers int, averageEngagementPercentage float64) float64 {
	base := math.Log2(float64(numFollowers))
	return base * averageEngagementPercentage
}

func numPossibleOrders(numPosts int) int {
	total := numPosts
	for i := total - 1; i > 0; i-- {
		total *= i
	}
	return total
}

func decayedFollowers(initialFollowers int, fractionLostDaily float64, days int) float64 {
	retentionRate := 1.0 - fractionLostDaily
	multiple := math.Pow(retentionRate, float64(days))
	return float64(initialFollowers) * multiple
}

func logScale(data []float64, base float64) []float64 {
	res := []float64{}
	base = math.Log(base)
	for _, num := range data {
		result := math.Log(num) / base
		res = append(res, math.Round(result))
	}
	return res
}

func averageFollowers(nums []int) float64 {
	if len(nums) == 0 {
		return 0
	}

	total := 0
	for _, num := range nums {
		total += num
	}

	return float64(total / len(nums))
}
