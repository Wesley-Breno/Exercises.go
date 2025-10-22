package main

import "fmt"

func main() {
	vector1 := []int{}
	vector2 := []int{}

	for i := 1; i <= 2; i++ {
		for ii := 1; ii <= 10; ii++ {
			var value int
			fmt.Print("Inform a value: ")
			fmt.Scan(&value)

			if i == 1 {
				vector1 = append(vector1, value)
			} else {
				vector2 = append(vector2, value)
			}
		}
	}

	finalVector := []int{}
	for i := 0; i < 10; i++ {
		finalVector = append(finalVector, vector1[i])
		finalVector = append(finalVector, vector2[i])
	}

	fmt.Println("Vector 1:", vector1)
	fmt.Println("Vector 2:", vector2)
	fmt.Println("Final vector:", finalVector)
}
