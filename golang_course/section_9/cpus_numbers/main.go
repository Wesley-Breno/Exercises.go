package main

import (
	"fmt"
	"runtime"
)

// Showing the total number of CPUs the machine has
func main() {
	fmt.Println(runtime.NumCPU())
}
