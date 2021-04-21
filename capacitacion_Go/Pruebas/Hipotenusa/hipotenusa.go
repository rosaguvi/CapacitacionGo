package main

import (
	"errors"
	"math"
)

func main() {

}

func calcularHipotenusa(opuesto, adyasente float32) (vlrhipotenusa float32, err error) {
	if opuesto <= 0 || adyasente <= 0 {
		err = errors.New("Error los lado deben ser mayores a cero, no se realizara el calculo")
	} else {
		vlrhipotenusa = float32(math.Pow((math.Pow(float64(opuesto), 2) + math.Pow(float64(adyasente), 2)), float64(0.5)))
	}
	return
}
