package main

import (
	"errors"
	"strconv"
)

func main() {

}

func validarNumero(numero int) (resultado string, err error) {
	if numero < 0 {
		err = errors.New("Error: Numero negativo no se procesara")
		return
	}
	numInv, _ := strconv.Atoi(invertirNumero(strconv.Itoa(numero)))
	resultado = "no palíndromo"
	if numInv == numero {
		resultado = "palíndromo"
		return
	}
	return
}

func invertirNumero(numero string) (nuevoNumero string) {
	for _, digito := range numero {
		nuevoNumero = string(digito) + nuevoNumero
	}
	return
}
