package main

import "fmt"

/*

* Arrays in Go are static and homogeneous. The data does not change type and if you defined that it has to have 10 positions, it will have 10 positions for its entire 'life'.

* If the elements are not defined at creation, the array will have its spaces filled with 0

 */

func main() {
	// Creating array of type float64 and which has only 3 elements
	var notas [3]float64
	fmt.Println(notas) // Elements not defined at creation, the array will have the 3 spaces filled with 0

	// Filling the values ​​of each space in the array
	notas[0], notas[1], notas[2] = 7.8, 4.3, 9.1
	fmt.Println(notas)

	// Getting the sum of the elements of an array
	total := 0.0
	for i := 0; i < len(notas); i++ {
		total += notas[i]
	}

	// Taking the average of the array and showing it on the screen
	media := total / float64(len(notas))
	fmt.Printf("Media: %.2f\n", media)
}
