package main

import "errors"

func main() {

}

func generarFibonacci(cnt int) (resultado int, err error) {
	if cnt < 0 {
		err = errors.New("Error: no se generaran terminos negativos ")
		return
	}
	resultado = valorFibonacci(cnt)
	return
}

func valorFibonacci(termino int) (valor int) {
	if termino == 0 {
		return 0
	}
	if termino == 1 {
		return 1
	}
	return valorFibonacci(termino-1) + valorFibonacci(termino-2)
}
