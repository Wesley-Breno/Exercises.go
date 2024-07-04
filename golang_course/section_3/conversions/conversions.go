package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Converting an integer to float64
	x := 2.4
	y := 2
	fmt.Println(x / float64(y))

	// Converting a float to an integer
	fmt.Println(int(x) / y)

	// Transforming a float64 into int
	nota := 6.9
	notaFinal := int(nota)
	fmt.Println(notaFinal)

	// This will not turn 97 into a string!!! It will return the value related to the ASCII table based on 97!!!
	fmt.Println("Teste " + string(97))

	// Transforming int into string
	fmt.Println("Teste " + strconv.Itoa(123)) // Use the 'strconv.Itoa(int)' package to convert the integer to a string

	// Transformando string em int
	num, _ := strconv.Atoi("123") // This function returns 2 values, the first is the number transformed into an integer, the second will be an error message if you enter the number incorrectly.
	fmt.Println(num - 122)

	// Transforming a string into boolean
	b, _ := strconv.ParseBool("true") // 'true' and '1' == True
	if b {                            // If the value of 'b' is True and managed to transform it into bool, it shows 'True' on the screen
		fmt.Println("Verdadeiro")
	}
}
