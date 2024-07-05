package main

import "fmt"

func main() {
	funcsESalarios := map[string]float64{
		"Jose josao":    1100.34,
		"Gabriel silva": 2500.23,
		"Pedro junior":  12312.46, // The last element must always have a comma
	}

	funcsESalarios["Wesley breno"] = 1350.0

	for nome, salario := range funcsESalarios {
		fmt.Printf("%s tem um salario de: %f\n", nome, salario)
	}
}
