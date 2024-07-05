package main

import "fmt"

func tipo(i interface{}) string {
	switch i.(type) {
	case int:
		return "inteiro"
	case float32, float64:
		return "real"
	case string:
		return "string"
	case func():
		return "funcao"
	default:
		return "Nao sei o tipo deste dado"
	}
}

func main() {
	fmt.Println(tipo("OI"))
	fmt.Println(tipo(12))
	fmt.Println(tipo(3.4))
}
