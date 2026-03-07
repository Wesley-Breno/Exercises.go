package main

import "fmt"

func main() {
	questions := []string{
		"Did you call the victim?",
		"Were you at the crime scene?",
		"Do you live near the victim?",
		"Did you owe the victim?",
		"Have you ever worked with the victim?",
	}

	var answers []int
	yesCount := 0

	for _, question := range questions {
		fmt.Printf("%s (0 - Yes, 1 - No): ", question)
		var answer int
		fmt.Scan(&answer)

		answers = append(answers, answer)
		if answer == 0 {
			yesCount++
		}
	}

	switch yesCount {
	case 2:
		fmt.Println("Suspect")
	case 3, 4:
		fmt.Println("Accomplice")
	case 5:
		fmt.Println("Murderer")
	default:
		fmt.Println("Innocent")
	}
}
