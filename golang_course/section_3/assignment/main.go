package main

import "fmt"

func main() {
	// Creation of 'complete' variable informing variable, its type and its value.
	var b byte = 3
	fmt.Println(b)

	// Short variable creation. Creates a variable with its value, without informing its type
	i := 3

	// Adding +3 to variable i (i = i + 3)
	i += 3

	// Subtracting -3 in variable i (i = i - 3)
	i -= 3

	// Dividing i by the value 2 (i = i / 2)
	i /= 2

	// Multiplying i by 2 (i = i * 2)
	i *= 2

	// Taking the remainder of the division of i / 2 (i = i % 2)
	i %= 2

	fmt.Println(i)

	// Defining variables in just one line
	x, y := 1, 2 // X = 1 e Y = 2
	fmt.Println(x, y)

	// Changing variable values
	x, y = y, x
	fmt.Println(x, y)
}
