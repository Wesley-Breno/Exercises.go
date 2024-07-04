package main

import "fmt"

func main() {
	// Write the 2 prints on a single line, even though they are separate prints.
	fmt.Print("Mesma")
	fmt.Print(" linha.")

	// Write the prints but make a line break
	fmt.Println(" Nova")
	fmt.Println("linha.")

	// Formatting a variable to show only the first 2 decimal numbers.
	x := 2.123456
	fmt.Printf("O valor de x é %.2f", x)

	// Formatting a float variable and showing its value
	y := 2.123456
	fmt.Printf("O valor de x é %f", y)

	// Showing different types of variables in a single print
	a := 1
	b := 1.9999
	c := false
	d := "opa"
	fmt.Printf("\n%d %f %.1f %t %s", a, b, b, c, d)

	// Using %v to not need to specify the variable type
	fmt.Printf("\n%v %v %v %v", a, b, c, d)

	/*
		\n | Line break
		%d | Informs that the variable is of type int
		%f | Informs that the variable is of type float
		%.2f | Informs that the variable is of float type and will only show its value and the first 2 decimal places
		%t | Informs that the variable is of type bool (true or false)
		%s | Informs that the variable is of type string
		%v | You will not need to inform the type of the variable, it will show regardless of the value
	*/
}
