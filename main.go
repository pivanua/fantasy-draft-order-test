package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {

	olympians := []string{
		"Belgium 🇧🇪",
		"France 🇫🇷",
		"Germany 🇩🇪",
		"Great Britain 🇬🇧",
		"Ireland 🇮🇪",
		"Israel 🇮🇱",
		"Mexico 🇲🇽",
		"Netherlands 🇳🇱",
		"Sweden 🇸🇪",
		"United States 🇺🇸",
	}
	fantasyFolk := []string{
		"Ayelex",
		"Clemens",
		"Colton",
		"Diego",
		"Jared",
		"Marshall",
		"Reevie",
		"Roy",
		"Shareeq",
		"Zec",
	}

	assignOlympianToFolk(olympians, fantasyFolk)

}

// randomly assign an olympian to one of the fantasyFolk
func assignOlympianToFolk(olympians []string, fantasyFolk []string) {
	fmt.Printf("\n🏈🏈🏈🏈🏈🏈🏈🏈🏈🏈🏈🏈🏈🏈🏈🏈🏈🏈🏈🏈🏈🏈🏈🏈🏈🏈🏈🏈🏈🏈\n\n")
	randOlympianIndexes := randRangeSlice(0, len(olympians))
	fmt.Printf("Randomized olympian indexes: %v\n", randOlympianIndexes)
	randFantasyIndexes := randRangeSlice(0, len(fantasyFolk))
	fmt.Printf("Randomized bro indexes:      %v\n", randFantasyIndexes)
	fmt.Printf("\n🏈🏈🏈🏈🏈🏈🏈🏈🏈🏈🏈🏈🏈🏈🏈🏈🏈🏈🏈🏈🏈🏈🏈🏈🏈🏈🏈🏈🏈🏈\n\n")
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
