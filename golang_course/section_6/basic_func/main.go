package main

import "fmt"

// Simple function, without parameters and return
func funcao1() {
	fmt.Println("Funcao 1")
}

// Function that receives 2 string parameters and displays them on the screen
func funcao2(param1 string, param2 string) {
	fmt.Printf("Funcao 2: param1: %s, param2: %s", param1, param2)
}

// Function that informs the value that is returned
func funcao3() string {
	return "Funcao 3"
}

// Function that receives 2 string parameters and returns a text informing them.
func funcao4(p1, p2 string) string {
	return fmt.Sprintf("Funcao 4: param1: %s, param2: %s", p1, p2)
}

// Function that informs that it will return 2 strings
func funcao5() (string, string) {
	return "Retorno 1", "Retorno 2"
}

func main() {
	funcao1()                         // Calling function1
	funcao2("f2_param1", "f2_param2") // Calling function2 and informing the 2 parameters

	// Calling function3 and function4 and storing the return in variables.
	r3, r4 := funcao3(), funcao4("f4_param1",
		"f4_param2")
	fmt.Println(r3)
	fmt.Println(r4)

	// Calling function5 and storing the return values ​​in variables
	str1, str2 := funcao5()
	fmt.Println(str1)
	fmt.Println(str2)
}
