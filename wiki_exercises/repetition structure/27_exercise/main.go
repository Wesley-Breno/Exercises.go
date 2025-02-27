package main

import "fmt"

func sum(values []int) int {
	total := 0
	for _, value := range values {
		total += value
	}
	return total
}

func main() {
	var numberOfClasses int
	fmt.Print("Enter the number of classes: ")
	fmt.Scan(&numberOfClasses)

	var studentsPerClass []int

	for {
		if numberOfClasses == 0 {
			break
		}

		var numberOfStudents int
		fmt.Printf("Enter the number of students in class %v: ", numberOfClasses)
		fmt.Scan(&numberOfStudents)

		if numberOfStudents > 40 {
			fmt.Println("The number of students in a class cannot exceed 40.")
			continue
		}

		numberOfClasses--
		studentsPerClass = append(studentsPerClass, numberOfStudents)
	}

	fmt.Println("The average number of students per class: ", sum(studentsPerClass)/len(studentsPerClass))
}
