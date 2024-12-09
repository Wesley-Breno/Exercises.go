package main

import "fmt"

func main() {
	for {
		var number int
		fmt.Print("Enter a number between 0 and 10: ")
		fmt.Scan(&number)

		if number > -1 && number < 11 {
			fmt.Println("You entered the number: ", number)
			break
		} else {
			fmt.Println("Invalid number, please enter the correct value.")
		}
	}
}
