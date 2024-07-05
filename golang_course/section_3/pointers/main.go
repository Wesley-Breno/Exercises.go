package main

import "fmt"

func main() {
	i := 1

	// Go doesn't have pointer arithmetic

	// Creating an int pointer with a null value
	var p *int = nil

	// Making 'p' store the memory address of 'i'
	p = &i

	// Taking value from pointer 'p' and incrementing +1
	*p++
	i++ // Does the same thing as above

	// Note that as 'i' and 'p' are pointers, it will show 3, as we are adding values ​​to shared variables. If one changes, they all change.
	fmt.Println(p, *p, i, &i)
}
