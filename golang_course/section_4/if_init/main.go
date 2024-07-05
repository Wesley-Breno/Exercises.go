package main

import (
	"fmt"
	"math/rand"
	"time"
)

func numeroAleatorio() int {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	return r.Intn(10)
}

func main() {
	if i := numeroAleatorio(); i > 5 {
		fmt.Println("Ganhou")
		fmt.Printf("'%v' dentro do if", i)
	} else {
		fmt.Println("Perdeu")
		fmt.Printf("%v dentro do else", i)
	}

	// fmt.Println(i) // It wouldn't work here
}
