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
