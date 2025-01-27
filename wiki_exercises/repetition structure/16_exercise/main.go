package main

import "fmt"

func main() {
	t1 := 0
	t2 := 1
	fmt.Printf("%v %v", t1, t2)
	for {
		t3 := t1 + t2
		fmt.Printf(" %v", t3)
		if t3 > 500 {
			break
		}
		t1 = t2
		t2 = t3
	}
}
