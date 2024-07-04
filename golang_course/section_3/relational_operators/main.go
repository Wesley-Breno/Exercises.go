package main

import (
	"fmt"
	"time"
)

func main() {
	// Is 'Banana' the same as 'Banana'? ('Banana' == 'Banana')
	fmt.Println("'Banana' é igual a 'Banana'?:", "Banana" == "Banana")

	// Is 3 different from 2? (3 != 2)
	fmt.Println("3 é diferente de 2?", 3 != 2)

	// Is 3 less than 2? (3 < 2)
	fmt.Println("3 é menor do que 2?", 3 < 2)

	// Is 3 greater than 2? (3 > 2)
	fmt.Println("3 é maior do que 2?", 3 > 2)

	// Is 3 less than or equal to 2? (3 <= 2)
	fmt.Println("3 é menor ou igual a 2?", 3 <= 2)

	// Is 3 greater than or equal to 2? (3 >= 2)
	fmt.Println("3 é maior ou igual 2?", 3 >= 2)

	//////////////////////////////////////////////

	// Comparing Unix times
	d1 := time.Unix(0, 0)
	d2 := time.Unix(0, 0)

	// Is date d1 the same as date d2?
	fmt.Println("A data d1 é igual a data d2?", d1 == d2)
	fmt.Println("A data d1 é igual a data d2?", d1.Equal(d2))

	/////////////////////////////////////////////

	// Comparing metadata from a struct
	type Pessoa struct {
		Nome string
	}

	p1 := Pessoa{"Joao"}
	p2 := Pessoa{"Joao"}

	// Is p1.Name the same as p2.Name?
	fmt.Println("O nome da pessoa p1 é igual ao nome da pessoa p2?", p1.Nome == p2.Nome)
}
