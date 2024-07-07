package main

import (
	"fmt"
	"strings"
)

// Struct person who receives first and last name
type pessoa struct {
	nome      string
	sobrenome string
}

// Person struct method that takes the person's full name by joining the attributes 'first name' and 'surname'.
func (p pessoa) getNomeCompleto() string {
	return p.nome + " " + p.sobrenome
}

// Method of the person struct that receives a new name to be changed. Separate by spaces and add the first name in the 'name' attribute and the second name in the 'surname' attribute.
func (p *pessoa) setNomeCompleto(nomeCompleto string) {
	partes := strings.Split(nomeCompleto, " ")
	p.nome = partes[0]
	p.sobrenome = partes[1]
}

func main() {
	// Defining person with their first and last name and showing the full name
	p1 := pessoa{"Pedro", "Silva"}
	fmt.Println(p1.getNomeCompleto())

	// Defining a new name using the 'set Full Name' method and then getting the full name that was changed.
	p1.setNomeCompleto("Maria Pereira")
	fmt.Println(p1.getNomeCompleto())
}
