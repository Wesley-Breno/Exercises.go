package main

import "fmt"

type Person struct {
	Age    int
	Height float64
}

func main() {
	agesHeights := []Person{
		{10, 1.40}, {11, 1.45}, {12, 1.50}, {13, 1.55}, {14, 1.60},
		{15, 1.65}, {16, 1.70}, {17, 1.75}, {18, 1.80},
		{19, 1.85}, {20, 1.90}, {21, 1.95}, {22, 2.00},
		{23, 2.10}, {24, 2.15}, {25, 2.20}, {26, 2.25},
		{27, 2.30}, {28, 2.35}, {29, 2.40}, {30, 2.45},
		{31, 2.50}, {32, 2.55}, {33, 2.60}, {34, 2.65},
		{35, 2.70}, {36, 2.75}, {37, 2.80}, {38, 2.85}, {39, 2.90},
	}

	averageHeight := 0.0
	studentsCount := 0

	for _, student := range agesHeights {
		averageHeight += student.Height
	}

	averageHeight /= float64(len(agesHeights))

	for _, student := range agesHeights {
		if student.Age > 13 && student.Height < averageHeight {
			studentsCount++
		}
	}

	fmt.Println("Number of students older than 13 with height below the average height")
	fmt.Println("Number of students:", studentsCount)
}
