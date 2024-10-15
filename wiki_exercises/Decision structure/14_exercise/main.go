package main

import "fmt"

func main() {
	var grade1 float64
	fmt.Print("Enter the student's first grade: ")
	fmt.Scan(&grade1)

	var grade2 float64
	fmt.Print("Enter the student's second grade: ")
	fmt.Scan(&grade2)

	average := (grade1 + grade2) / 2

	var letterGrade string
	if average <= 4 {
		letterGrade = "E"
	} else if average <= 6 {
		letterGrade = "D"
	} else if average <= 7.5 {
		letterGrade = "C"
	} else if average <= 9 {
		letterGrade = "B"
	} else {
		letterGrade = "A"
	}

	fmt.Println("Student's grades:", grade1, grade2)
	fmt.Println("Student's average:", average)
	fmt.Println("Letter grade:", letterGrade)

	fmt.Print("Status: ")
	if letterGrade == "A" || letterGrade == "B" || letterGrade == "C" {
		fmt.Println("Approved")
	} else {
		fmt.Println("Failed")
	}
}
