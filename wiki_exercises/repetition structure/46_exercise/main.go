package main

import (
	"fmt"
	"slices"
)

func SumFloat64(slice []float64) float64 {
	result := 0.0
	for _, value := range slice {
		result += value
	}
	return result
}

type Athlete struct {
	Name        string
	Jumps       []float64
	BestJump    float64
	WorstJump   float64
	AverageJump float64
}

func main() {
	var athletes = make(map[int]Athlete)
	counter := 1

	for {
		var athleteName string
		fmt.Print("Enter the athlete's name: ")
		fmt.Scanln(&athleteName)

		if athleteName == "" {
			break
		}

		var formattedJumps []float64

		for i := 0; i < 5; i++ {
			var jump float64
			fmt.Printf("\nEnter the distance of jump #%v: ", i+1)
			fmt.Scan(&jump)
			formattedJumps = append(formattedJumps, jump)
		}

		bestJump := slices.Max(formattedJumps)
		worstJump := slices.Min(formattedJumps)
		var jumpsWithoutBestAndWorst []float64

		skippedBest := false
		skippedWorst := false
		for _, jump := range formattedJumps {
			if jump == bestJump && !skippedBest {
				skippedBest = true
				continue
			}
			if jump == worstJump && !skippedWorst {
				skippedWorst = true
				continue
			}
			jumpsWithoutBestAndWorst = append(jumpsWithoutBestAndWorst, jump)
		}

		averageJump := SumFloat64(jumpsWithoutBestAndWorst) / float64(len(jumpsWithoutBestAndWorst))

		athletes[counter] = Athlete{
			Name:        athleteName,
			Jumps:       formattedJumps,
			BestJump:    bestJump,
			WorstJump:   worstJump,
			AverageJump: averageJump,
		}

		counter++
	}

	if len(athletes) > 0 {
		athleteScores := make(map[string][]float64)

		for id, data := range athletes {
			fmt.Println()
			fmt.Println("---------------------------------")
			fmt.Println("Athlete ID:", id)
			fmt.Println("Best jump:", data.BestJump)
			fmt.Println("Worst jump:", data.WorstJump)
			fmt.Println("Average of remaining jumps:", data.AverageJump)
			fmt.Println("Athlete name:", data.Name)
			fmt.Println("Jumps:", data.Jumps)

			athleteScores[data.Name] = data.Jumps
		}

		fmt.Println("\nFinal Results:")
		for name, jumps := range athleteScores {
			fmt.Printf("%s: %v m\n", name, jumps)
		}
	}
}
