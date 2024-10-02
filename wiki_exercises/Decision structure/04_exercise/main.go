package main

import (
	"fmt"
	"strings"
)

func main() {
	var letter string
	fmt.Print("Enter a letter: ")
	fmt.Scanln(&letter)

	if !strings.ContainsAny(letter, "aeiou") && strings.ContainsAny(letter, "bcdfgjklmnpqrstvwxyz") {
		fmt.Println("The letter is a consonant.")
	} else if strings.ContainsAny(letter, "aeiou") && !strings.ContainsAny(letter, "bcdfgjklmnpqrstvwxyz") {
		fmt.Println("The letter is a vowel")
	} else {
		fmt.Println("Please enter only one letter of the alphabet.")
	}
}
