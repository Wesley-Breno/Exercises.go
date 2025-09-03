package main

import "fmt"

func main() {
	consonants := []string{}
	for i := 0; i < 10; i++ {
		var letter string
		fmt.Println()
		fmt.Print("Enter a letter: ")
		fmt.Scan(&letter)

		if letter == "a" || letter == "e" || letter == "i" || letter == "o" || letter == "u" {
			continue
		}

		consonants = append(consonants, letter)
	}

	fmt.Println("\nTyped consonants:", consonants)
	fmt.Println("Number of consonants:", len(consonants))
}
