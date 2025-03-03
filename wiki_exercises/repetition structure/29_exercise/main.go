package main

import "fmt"

func main() {
	fmt.Println("Lojas Quase Dois - Tabela de preços")
	for i := 1; i < 51; i++ {
		fmt.Printf("[%v] = R$ %v\n", i, float64(i)*1.99)
	}
}
