package main

import (
	"errors"
	"math"
)

func main() {

}

func calcularPerimetro(radio float32) (perimetro float32, err error) {
	if radio < 0 {
		err = errors.New("No se realiza calculo con valores negativos")
	} else {
		perimetro = 2 * math.Pi * radio
	}
	return
}
func calcularArea(radio float32) (area float32, err error) {
	if radio < 0 {
		err = errors.New("No se realiza calculo con valores negativos")
	} else {
		area = float32(float64(math.Pi) * math.Pow(float64(radio), 2))
	}
	return
}
