package main

import "fmt"

func main() {
	var side1 int
	fmt.Print("Enter the 1st side of the triangle: ")
	fmt.Scanln(&side1)

	var side2 int
	fmt.Print("Enter the 2nd side of the triangle: ")
	fmt.Scanln(&side2)

	var side3 int
	fmt.Print("Enter the 3rd side of the triangle: ")
	fmt.Scanln(&side3)

	if side1 < side2+side3 && side2 < side1+side3 && side3 < side2+side1 {
		if side1 == side2 && side1 == side3 {
			fmt.Println("Equilateral")
		} else if side1 == side2 || side1 == side3 || side2 == side3 {
			fmt.Println("Isosceles")
		} else {
			fmt.Println("Scalene")
		}
	} else {
		fmt.Println("The given values do not form a triangle")
	}
}
