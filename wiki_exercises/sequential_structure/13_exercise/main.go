package main

import "fmt"

func main() {
	var height float64
	fmt.Print("Enter your height: ")
	fmt.Scanln(&height)

	var sex string
	fmt.Print("Enter your gender using 'F' or 'M': ")
	fmt.Scanln(&sex)

	if sex == "F" {
		fmt.Println("Your ideal weight is:", (62.1*height)-44.7)
	}

	if sex == "M" {
		fmt.Println("Your ideal weight is:", (72.7*height)-58)
	}

	if sex != "F" && sex != "M" {
		fmt.Println("Please enter your data correctly")
	}

}
