package main

import "math"

// Data that starts with a capital letter is considered 'public', that is, it is possible to see and use it in other Go code files.

// Data that starts with a lowercase letter is private, that is, it is only visible and accessible within its own package.

// Using a period in nomenclatures will generate something public

// Using point or _Point in nomenclatures will generate something private

type Ponto struct {
	x float64
	y float64
}

func catetos(a, b Ponto) (cx, cy float64) {
	cx = math.Abs(b.x - a.x)
	cy = math.Abs(b.y - a.y)
	return
}

// Distance is responsible for calculating the linear distance between two points
func Distancia(a, b Ponto) float64 {
	cx, cy := catetos(a, b)
	return math.Sqrt(math.Pow(cx, 2) + math.Pow(cy, 2))
}
