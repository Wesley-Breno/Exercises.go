package main

import "fmt"

func main() {
	var grade1 float64
	fmt.Print("Enter your first note: ")
	fmt.Scanln(&grade1)

	var grade2 float64
	fmt.Print("Enter your second note: ")
	fmt.Scanln(&grade2)

	average := (grade1 + grade2) / 2

	if average == 10 {
		fmt.Println("Approved with distinction!")
	} else if average >= 7 {
		fmt.Println("Approved!")
	} else {
		fmt.Println("Failed ;(")
	}
}
