package main

import "fmt"

func main() {
	questions := []string{
		"Did you call the victim? ",
		"Were you at the crime scene? ",
		"Do you live near the victim? ",
		"Did you owe the victim? ",
		"Have you worked with the victim? ",
	}

	fmt.Println("-----------------------")
	fmt.Println("Answer with 0 for no")
	fmt.Println("Answer with 1 for yes")
	fmt.Println("-----------------------")

	var result int
	for _, question := range questions {
		var num int
		fmt.Println()
		fmt.Print(question)
		fmt.Scan(&num)
		if num > 0 {
			result++
		}
	}

	if result == 2 {
		fmt.Println("\nYou are a suspect o.O")
	} else if result == 3 || result == 4 {
		fmt.Println("\nYou are an accomplice O.O")
	} else if result == 5 {
		fmt.Println("\nYou are the murderer ＞﹏＜")
	} else {
		fmt.Println("\nYou are innocent \\o/")
	}
}
