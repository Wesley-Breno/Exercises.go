package main

import "fmt"

func main() {
	var vector []int

	for i := 0; i < 5; i++ {
		var num int
		fmt.Print("Enter a number: ")
		fmt.Scan(&num)

		vector = append(vector, num)
	}

	fmt.Println("The numbers you entered are:")
	for i, num := range vector {
		fmt.Printf("Number %dª: %d\n", i+1, num)
	}
}