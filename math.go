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
	for i := 0; i < numMonths; i++ {
		switch influencerType {
		case "fitness":
			followerCount *= 4
		case "cosmetic":
			followerCount *= 3
		default:
			followerCount *= 2
		}
	}
	return followerCount
}
