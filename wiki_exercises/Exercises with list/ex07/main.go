package main

import "fmt"

func main() {
	nums := []int{}
	soma := 0
	multi := 1

	for i := 0; i < 5; i++ {
		var num int
		fmt.Print("Inform a number: ")
		fmt.Scan(&num)
		nums = append(nums, num)
		fmt.Println()
	}

	for _, num := range nums {
		soma += num
		multi *= num
	}

	fmt.Println("Sum:", soma)
	fmt.Println("Multiplication:", multi)
	fmt.Println("Numbers:", nums)
}
