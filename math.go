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
