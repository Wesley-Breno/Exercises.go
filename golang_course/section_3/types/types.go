package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	// [[[[ INT NUMBERS ]]]]
	fmt.Println(1, 2, 1000)

	// Showing the type of a data
	fmt.Println("Literal inteiro é", reflect.TypeOf(32000))

	// no sign (only positive)... uint8, uint16, uint32 and uint64
	var b byte = 255
	fmt.Println("O byte é", reflect.TypeOf(b))

	// with sign... int8, int16, int32 and int64
	i1 := math.MaxInt64 // Get the largest integer value up to 64 bits
	fmt.Println("O valor maximo do int é", i1)
	fmt.Println("O tipo de i1 é", reflect.TypeOf(i1))

	var i2 rune = 'a' // represents a mapping of the Unicode table (int32)
	fmt.Println("O rune é", reflect.TypeOf(i2))
	fmt.Println(i2)

	// [[[[ Float Numbers (float32, float64) ]]]]
	var x float32 = 49.99
	fmt.Println("O tipo de x é", reflect.TypeOf(x))
	fmt.Println("O tipoe do literal 49.99 é", reflect.TypeOf(49.99)) // float64

	// [[[[ BOOLEANS ]]]]
	bo := true
	fmt.Println("O tipo de bo é", reflect.TypeOf(bo))
	fmt.Println(!bo) // Shows the opposite value of 'bo' (only works if it is of type 'bool')

	// [[[[ STRINGS ]]]]
	s1 := "Ola meu nome é Leo"
	fmt.Println(s1 + "!")
	fmt.Println("O tamanho da string é", len(s1))

	// String with multiple lines
	s2 := `Ola
	meu
	nome
	é
	Leo`
	fmt.Println("O tamanho da string é", len(s2))

}
