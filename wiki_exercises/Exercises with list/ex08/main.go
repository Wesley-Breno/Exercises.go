package main

import "fmt"

func reverseString(s string) string {
    runes := []rune(s)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}

func main() {
	people := [][]string{}

	for i := 0; i < 5; i++ {
		var age int
		height := 0.0
		fmt.Printf("Enter the age of person %d: ", i+1)
		fmt.Scan(&age)
		fmt.Print("Enter the height: ")
		fmt.Scan(&height)
		people = append(people, []string{fmt.Sprintf("%d", age), fmt.Sprintf("%.2f", height)})
		fmt.Println()
	}

	for _, person := range people {
		fmt.Printf("Age: %s\n", reverseString(person[0]))
		fmt.Printf("Height: %s\n", reverseString(person[1]))
	}
}
