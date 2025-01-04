package main

import "fmt"

func main() {
	for i := 1; i < 21; i++ {
		fmt.Println(i)
	}

	fmt.Println()

	var cont int64
	for i := 1; i < 21; i++ {
		fmt.Printf("%v ", i)
		cont++
		if cont >= 2 {
			fmt.Println()
			cont = 0
		}
	}
}
