package main

import "fmt"

func notaParaConceito(n float64) string {
	var nota = int(n)
	switch nota {
	case 10:
		fallthrough // Make him continue to the next cases. It's like 'continue' in Python
	case 9: // If 'nota' is 9
		return "A"
	case 8, 7: // If the 'nota' is 8 or 7
		return "B"
	case 6, 5:
		return "C"
	case 4, 3:
		return "D"
	case 2, 1, 0:
		return "E"
	default: // If none of the above cases
		return "Nota invalida"
	}
}

func main() {
	fmt.Println(notaParaConceito(9.8))
	fmt.Println(notaParaConceito(6.9))
	fmt.Println(notaParaConceito(2.1))
	fmt.Println(notaParaConceito(-1))
}
