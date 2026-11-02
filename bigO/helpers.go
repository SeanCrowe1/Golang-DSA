package big_o

import (
	"fmt"
	"math/rand"
)

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
