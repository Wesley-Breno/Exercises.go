package main

import (
	"fmt"
	"math/rand"
	"time"
)

func shuffleString(s string) string {
	runes := []rune(s)

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(runes), func(i, j int) {
		runes[i], runes[j] = runes[j], runes[i]
	})

	return string(runes)
}

func main() {
	wesley := "wesley"
	shuffledWesley := shuffleString(wesley)
	fmt.Println("Original:", wesley)
	fmt.Println("Shuffled:", shuffledWesley)
	
	fmt.Println()
	
	jose := "jose"
	shuffledJose := shuffleString(jose)	
	fmt.Println("Original:", jose)
	fmt.Println("Shuffled:", shuffledJose)
}
