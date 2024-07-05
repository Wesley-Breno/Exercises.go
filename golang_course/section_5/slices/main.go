package main

import (
	"fmt"
	"reflect"
)

func main() {
	a1 := [3]int{1, 2, 3} // array
	s1 := []int{1, 2, 3}  // slice
	fmt.Println(a1, s1)

	fmt.Println(reflect.TypeOf(a1), reflect.TypeOf(s1))

	a2 := [5]int{1, 2, 3, 4, 5} // array

	// slice is not an array! Slice defines a piece of an array
	s2 := a2[1:3] // slice
	fmt.Println(a2, s2)

	s3 := a2[:2] // New slice, but points to the same array

	s4 := s1[:1] // The slice of a slice
	fmt.Println(s3, s4)
}
