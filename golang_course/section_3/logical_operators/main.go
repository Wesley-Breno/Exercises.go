package main

import "fmt"

func compras(trab1, trab2 bool) (bool, bool, bool) {
	comprarTv50 := trab1 && trab2    // Worked twice, bought 50 inch TV
	comprarTv32 := trab1 != trab2    // Only worked once, bought 32 inch TV
	comprarSorvete := trab1 || trab2 // If you worked, buy ice cream

	return comprarTv50, comprarTv32, comprarSorvete
}

func main() {
	tv50, tv32, sorvete := compras(true, true)
	fmt.Printf("Tv50: %t, Tv32: %t, Sorvete: %t, Saudavel: %t", tv50, tv32, sorvete, !sorvete)
}
