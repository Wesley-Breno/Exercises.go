package main

import (
	"fmt"
	"unicode"
)

func onlyNumbers(s string) bool {
	for _, char := range s {
		if !unicode.IsDigit(char) {
			return false
		}
	}
	return len(s) > 0
}

func main() {
	var number string
	fmt.Print("Enter a number less than 1000: ")
	fmt.Scan(&number)

	if len(number) < 4 && onlyNumbers(number) == true {
		if len(number) == 3 {
			fmt.Printf("%v hundred", string(number[0]))
			firstChar := int(number[0] - '0')
			if int(firstChar) > 1 {
				fmt.Print("s")
			}

			fmt.Printf(", %v ten", string(number[1]))
			secondChar := int(number[1] - '0')
			if int(secondChar) > 1 {
				fmt.Print("s")
			}

			fmt.Printf(" and %v unit", string(number[2]))
			lastChar := int(number[2] - '0')
			if int(lastChar) > 1 {
				fmt.Print("s")
			}
		} else if len(number) == 2 {
			fmt.Printf("%v ten", string(number[0]))
			firstChar := int(number[0] - '0')
			if int(firstChar) > 1 {
				fmt.Print("s")
			}

			fmt.Printf(" and %v unit", string(number[1]))
			lastChar := int(number[1] - '0')
			if int(lastChar) > 1 || int(lastChar) == 0 {
				fmt.Print("s")
			}
		} else {
			fmt.Printf("%v unit", string(number[0]))
			lastChar := int(number[0] - '0')
			if int(lastChar) > 1 || int(lastChar) == 0 {
				fmt.Print("s")
			}
		}
	} else {
		fmt.Println("Please enter a valid number!")
	}
}
