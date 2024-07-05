package main

import (
	"fmt"
	"time"
)

func main() {
	// Counting from 1 to 10
	i := 1
	for i <= 10 { // There is no while loop in Go
		fmt.Println(i)
		i++
	}

	// Count from 0 to 20 by jumping 2 by 2
	for i := 0; i <= 20; i += 2 {
		fmt.Printf("%d ", i)
	}

	// Counting from 1 to 10 and saying which number is even and odd
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Printf("\n%d é Par", i)
		} else {
			fmt.Printf("\n%d é Impar", i)
		}
	}

	// infinite loop printing 'Para sempre...'
	for {
		fmt.Println("Para sempre...")
		time.Sleep(time.Second) // Waiting for 1 second
	}
}
