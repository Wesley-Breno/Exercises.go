package main

import "fmt"

func main() {
	var sex string
	fmt.Print("Enter the gender with 'F' or 'M': ")
	fmt.Scanln(&sex)

	if sex == "F" {
		fmt.Println("F - Female")
	} else if sex == "M" {
		fmt.Println("M - Male")
	} else {
		fmt.Println("Invalid gender.")
	}
}
