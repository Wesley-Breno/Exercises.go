package main

import "fmt"

func main() {
	var number int
	fmt.Print("Enter a number: ")
	fmt.Scanln(&number)

	if number > 0 {
		fmt.Println("You entered a positive number")
	} else {
		fmt.Println("You entered a negative number")
	}
}
