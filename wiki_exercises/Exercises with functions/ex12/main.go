package main

import (
	"fmt"
)

func drawRectangle(width, height int) {
	signs := []string{"+", "-", "|"}

	for {
		if width >= height {
			signIndex := 0

			// Top border
			for i := 0; i != width; i++ {
				fmt.Print(signs[signIndex])
				signIndex++
				if signIndex > 2 {
					signIndex = 0
				}
			}
			fmt.Println()

			// Middle part
			for i := 0; i != height; i++ {
				fmt.Print(signs[signIndex])
				signIndex++
				if signIndex > 2 {
					signIndex = 0
				}
				for j := 0; j != width-2; j++ {
					fmt.Print(" ")
				}
				fmt.Print(signs[signIndex])
				signIndex++
				if signIndex > 2 {
					signIndex = 0
				}
				fmt.Println()
			}

			// Bottom border
			for i := 0; i != width; i++ {
				fmt.Print(signs[signIndex])
				signIndex++
				if signIndex > 2 {
					signIndex = 0
				}
			}
			break
		} else {
			width = 10
			height = 10
		}
	}
}

func main() {
	drawRectangle(100, 100)
}
