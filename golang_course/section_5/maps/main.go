package main

import "fmt"

func main() {
	// Maps must be initialized
	aprovados := make(map[int]string)

	aprovados[1] = "Maria"
	aprovados[2] = "Pedro"
	aprovados[3] = "Ana"

	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF: %d)\n", nome, cpf)
	}

	fmt.Println(aprovados[3])
	delete(aprovados, 3)
	fmt.Println(aprovados[3])
}
