package main

import "fmt"

func main() {
	var num int
	fmt.Print("Enter the number: ")
	fmt.Scan(&num)

	t1 := 0
	t2 := 1
	fmt.Printf("%v %v", t1, t2)
	for i := 1; i < num; i++ {
		t3 := t1 + t2
		fmt.Printf(" %v", t3)
		t1 = t2
		t2 = t3
	}
}
