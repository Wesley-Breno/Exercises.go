package main

import (
	"fmt"
)

type AthleteData struct {
	HighestScore float64
	LowestScore  float64
	Average      float64
}

func GetHighestScore(scores []float64) float64 {
	highest := 0.0
	for _, s := range scores {
		if s > highest {
			highest = s
		}
	}
	return highest
}

func GetLowestScore(scores []float64) float64 {
	lowest := 0.0
	for i, s := range scores {
		if i == 0 {
			lowest = s
			continue
		}
		if s < lowest {
			lowest = s
		}
	}
	return lowest
}

func GetSum(nums []float64) float64 {
	sum := 0.0
	for _, n := range nums {
		sum += n
	}
	return sum
}

func main() {
	athletes := map[string][]float64{}

	fmt.Println("\n[Enter '-1' as the athlete's name to end the program]")

	for {
		var athleteName string = ""
		fmt.Print("\nEnter the athlete's name: ")
		fmt.Scan(&athleteName)

		if athleteName == "-1" {
			break
		}

		for i := 1; i <= 7; i++ {
			var score float64
			fmt.Printf("\nEnter the score given by judge #%v: ", i)
			fmt.Scan(&score)
			athletes[athleteName] = append(athletes[athleteName], score)
		}
	}

	if len(athletes) > 0 {
		for name, scores := range athletes {
			highest := GetHighestScore(scores)
			lowest := GetLowestScore(scores)

			var filteredScores = []float64{}
			for _, s := range scores {
				if s != highest && s != lowest {
					filteredScores = append(filteredScores, s)
				}
			}
			sum := GetSum(filteredScores)
			average := sum / float64(len(filteredScores))

			fmt.Println("\n\n[ Final Result ]\n")
			fmt.Println("Athlete: ", name)
			fmt.Println("Highest score: ", highest)
			fmt.Println("Lowest score: ", lowest)
			fmt.Println("Average: ", average)
		}
	} else {
		fmt.Println("You didn't enter any athlete")
	}
}
