package main

import "fmt"

func calculateSum(grades []float64) float64 {
	var sum float64
	for _, grade := range grades {
		sum += grade
	}

	return sum
}

func calculateAverage(grades []float64) float64 {
	if len(grades) == 0 {
		return 0
	}

	return calculateSum(grades) / float64(len(grades))
}

func main() {
	var grades []float64
	var count = 1

	for {
		var grade float64
		fmt.Printf("Enter grade %v: ", count)
		fmt.Scan(&grade)

		if grade == -1 {
			break
		}

		grades = append(grades, grade)
		count++
	}

	if len(grades) == 0 {
		fmt.Println("No grades were provided.")
		fmt.Println("Program finished.")
		return
	}

	aboveAverageCount := 0
	belowSevenCount := 0
	sum := calculateSum(grades)
	average := calculateAverage(grades)

	for _, grade := range grades {
		if grade > average {
			aboveAverageCount++
		}
		if grade < 7 {
			belowSevenCount++
		}
	}

	fmt.Println("Result")
	fmt.Printf("Number of grades read: %d\n", len(grades))
	fmt.Print("Grades in input order: ")
	for _, grade := range grades {
		fmt.Printf("%.2f ", grade)
	}
	fmt.Println()
	fmt.Println("Grades in reverse order:")
	for i := len(grades) - 1; i >= 0; i-- {
		fmt.Printf("%.2f\n", grades[i])
	}
	fmt.Printf("Sum of grades: %.2f\n", sum)
	fmt.Printf("\nAverage: %.2f\n", average)
	fmt.Printf("Number of grades above average: %v\n", aboveAverageCount)
	fmt.Printf("Number of grades below seven: %v\n", belowSevenCount)
	fmt.Println("Program finished.")
}
