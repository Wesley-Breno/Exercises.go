package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()

	// switch true, will look for the first case that is true
	switch {
	case t.Hour() < 12:
		fmt.Println("Bom dia!")
	case t.Hour() < 18:
		fmt.Println("Boa tarde!")
	default:
		fmt.Println("Boa noite.")
	}
}
