package main

import "fmt"

func main() {
	sumGrades := 0.0
	for i := 1; i <= 3; i++ {
		var grade float64
		fmt.Printf("\nEnter the %vº grade of the student: ", i)
		fmt.Scan(&grade)
		sumGrades += grade
	}

	average := sumGrades / 3

	if average == 10 {
		fmt.Println("Passed with Distinction")
	} else if average >= 7 {
		fmt.Println("Passed")
	} else {
		fmt.Println("Failed")
	}
}
