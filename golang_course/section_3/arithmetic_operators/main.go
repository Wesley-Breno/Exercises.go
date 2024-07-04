package main

import (
	"fmt"
	"math"
)

func main() {
	a := 3
	b := 2

	// Basic operators
	fmt.Println("Soma =>", a+b)
	fmt.Println("Subtracao =>", a-b)
	fmt.Println("Divisao =>", a/b)
	fmt.Println("Multiplicacao =>", a*b)
	fmt.Println("Modulo (Resto da divisao) =>", a%b)

	// Bitwise operations
	fmt.Println("E =>", a&b)
	fmt.Println("OU =>", a|b)
	fmt.Println("Xor =>", a^b)

	// Operations using the math package...

	c := 3.0
	d := 2.0

	// Largest number between A and B
	fmt.Println("Maior numero entre A e B =>", math.Max(float64(a), float64(b))) // Most math methods consider parameters to be float64, remember to switch them to float64 if they are not.

	// Smallest number between C and D
	fmt.Println("Menor valor entre C e D =>", math.Min(c, d)) // In this case, C and D are already float64

	// C raised to D
	fmt.Println("Exponenciacao C elevado a D =>", math.Pow(c, d))
}
