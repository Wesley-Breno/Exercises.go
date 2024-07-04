package main

import "fmt"

// Function that adds 2 numbers
func somar(a int, b int) int {
	return a + b
}

// Function that prints an integer value on the screen
func imprimir(valor int) {
	fmt.Println(valor)
}

// Main function (algorithm routine)
func main() {
	resultado := somar(3, 4)
	imprimir(resultado)
}
