package main

import "fmt"

// Function that multiplies the values ​​of 2 parameters
func multiplicacao(n1, n2 int) int {
	return n1 * n2
}

/*
Function that takes a function as a parameter, executes it and returns the result

In this case, the function that is received as a parameter is 'multiplication', it is executed and its return is passed.
*/
func exec(funcao func(int, int) int, p1, p2 int) int {
	return funcao(p1, p2)
}

func main() {
	resultado := exec(multiplicacao, 3, 4)
	fmt.Println(resultado)
}
