package main

import "fmt"

func getMaxValue(students map[int]int) map[int]int {
	var maxHeight int
	var studentID int
	var studentResponse = make(map[int]int)

	for id, height := range students {
		if height > maxHeight {
			studentID = id
			maxHeight = height
		}
	}
	studentResponse[studentID] = maxHeight
	return studentResponse
}

func getMinValue(students map[int]int) map[int]int {
	var minHeight int
	var studentID int
	var studentResponse = make(map[int]int)

	for id, height := range students {
		if height < minHeight || minHeight == 0 && studentID == 0 {
			studentID = id
			minHeight = height
		}
	}
	studentResponse[studentID] = minHeight
	return studentResponse
}

func main() {
	var students = make(map[int]int)

	for i := 0; i < 10; i++ {
		var id int
		var heightCm int

		fmt.Print("\nEnter student ID: ")
		fmt.Scan(&id)

		fmt.Print("\nEnter student height: ")
		fmt.Scan(&heightCm)

		fmt.Println("\n-")

		students[id] = heightCm
	}

	fmt.Println("Showing...")
	shortestStudent := getMinValue(students)
	tallestStudent := getMaxValue(students)
	fmt.Println("Shortest student")
	for id, height := range shortestStudent {
		fmt.Printf("ID %v\nHeight CM: %v", id, height)
	}

	fmt.Println("\n\nTallest student")
	for id, height := range tallestStudent {
		fmt.Printf("ID %v\nHeight CM: %v", id, height)
	}
}
