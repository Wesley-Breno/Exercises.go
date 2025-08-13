package utils

import (
	"strconv"
)

func ElementsIsDigit(elements []string) bool {
	var valores []string
    for _, valor := range elements {
        if _, err := strconv.Atoi(valor); err == nil {
            valores = append(valores, valor)
        }
    }
	return len(valores) == len(elements)
}