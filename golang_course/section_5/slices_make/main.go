package main

import "fmt"

func main() {
	// slice of 10 spaces using make() method
	s := make([]int, 10)
	s[9] = 12 // space 9 of the slices will be 12 now
	fmt.Println(s)

	// Array that can support up to 20 elements, but now it will only occupy 10 spaces
	s = make([]int, 10, 20)
	fmt.Println(s, len(s), cap(s))

	// Adding 10 more elements to the slice (reached the supported capacity of 20 elements)
	s = append(s, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0)
	fmt.Println(s, len(s), cap(s))

	// Trying to add 1 more element after reaching the slice's supported limit
	s = append(s, 1)
	fmt.Println(s, len(s), cap(s))
}
