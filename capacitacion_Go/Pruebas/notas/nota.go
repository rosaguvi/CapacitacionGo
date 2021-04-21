package main

import (
	"errors"
	"fmt"
)

func main() {

}

func calcularNota(notas []float32) (nota float32, err error) {
	if len(notas) != 3 {
		err = errors.New("La cantidad de notas no es la adecuada, no se realizara la operación")
	} else {
		notaAprueba := float32(3)
		totalNotas := float32(0)
		for _, valor := range notas {
			if valor > 5 {
				notaAprueba = 6
			}
			if valor < 0 {
				err = errors.New("Error: una nota no puede ser menor acero, no se realizara la operación")
				totalNotas = 0
				break
			}
			totalNotas += valor
		}
		if (totalNotas * 0.25) >= notaAprueba {
			fmt.Println("estoy aqui...", notas, totalNotas, notaAprueba)
			nota = 0
		} else {
			nota = (float32(notaAprueba) - (totalNotas * 0.25)) / 0.25
			if (notaAprueba == 3 && nota > 5) || (notaAprueba == 6 && nota > 10) {
				err = errors.New("Error: ya no es posible aprobar la materia")
			}
		}
	}
	return
}
