package main

import (
	"fmt"
	"math/rand/v2"
)

// randomly assign an olympian to one of the fantasyFolk
func assignOlympianToFolk(olympians []string, fantasyFolk []string) {
	randOlympianIndexes := randRangeSlice(0, len(olympians))
	fmt.Printf("Randomized olympian indexes: %v\n", randOlympianIndexes)
	randFantasyIndexes := randRangeSlice(0, len(fantasyFolk))
	fmt.Printf("Randomized ppl indexes:      %v\n", randFantasyIndexes)
	defer println("\n...all done, bye.")

	for i, olympianIndex := range randOlympianIndexes {
		olympian := olympians[olympianIndex]
		bro := fantasyFolk[randFantasyIndexes[i]]
		fmt.Printf("%s ===> %s \n", bro, olympian)
	}
}

func randRangeSlice(min, max int) []int {
	randSlice := make([]int, max-min)
	for i := min; i < max; i++ {
		randSlice[i] = i
	}
	// shuffle shuffle shuffle
	rand.Shuffle(len(randSlice), func(i, j int) {
		randSlice[i], randSlice[j] = randSlice[j], randSlice[i]
	})
	return randSlice
}
