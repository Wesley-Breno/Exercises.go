package main

import (
	"fmt"
	m "math" // Setting a nickname for the package
)

func main() {
	const PI float64 = 3.1415 // constant of type float64
	var raio = 3.2            // Type (float64) inferred by the compiler

	// short way to create a var
	area := PI * m.Pow(raio, 2)
	fmt.Println("A area da circunferencia é", area)

	// Constants building block
	const (
		a = 1
		b = 2
	)

	// Variable building block
	var (
		c = 3
		d = 4
	)
	fmt.Println(a, b, c, d)

	// Defining 2 variables in a single line
	var e, f bool = true, false
	fmt.Println(e, f)

	// Defining 3 variables in reduced form in a single line
	g, h, i := 2, false, "epa!"
	fmt.Println(g, h, i)
}
