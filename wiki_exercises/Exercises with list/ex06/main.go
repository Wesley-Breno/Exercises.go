package main

import "fmt"

func main() {
	alunosNaMedia := []float64{}

	for i := 0; i < 10; i++ {
		notas := []float64{}

		for ii := 0 ; ii < 4; ii++ {
			var number float64
			fmt.Printf("Enter the grade %d of student %d: ", ii+1, i+1)
			fmt.Scan(&number)
			notas = append(notas, number)
		}

		media := 0.0
		for _, nota := range notas {
			media += nota
		}
		media /= float64(len(notas))
		if media >= 7 {
			alunosNaMedia = append(alunosNaMedia, media)
		}
	}

	fmt.Println("Students with average grades >= 7:", len(alunosNaMedia))
}
