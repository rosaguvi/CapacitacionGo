package main

import (
	"errors"
	"math"
)

func main() {

}
func validarNumeros(numeroval int) (resul int, err error) {
	if numeroval < 0 {
		err = errors.New("Nuermo Negativo no se validara...")
	} else {
		resul = int(math.Mod(float64(numeroval), 2))
	}
	return
}
