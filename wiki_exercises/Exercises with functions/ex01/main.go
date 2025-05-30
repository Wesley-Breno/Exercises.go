package main

import "fmt"

func PrintTreeOfNValues(n int) {
	count := 1
	for {
		for i := 0; i < count; i++ {
			fmt.Print(count)
		}
		fmt.Println()

		if count == n {
			break
		}
		count++
	}
}

func main() {
	PrintTreeOfNValues(9)
}
