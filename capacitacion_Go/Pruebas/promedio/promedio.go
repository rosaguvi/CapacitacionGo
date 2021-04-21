package main

import "errors"

func main() {

}

func calcularPromedio(numeros []float32) (promedio float32, err error) {
	promedio = float32(0)
	if len(numeros) == 0 {
		err = errors.New("no hay numeros para generar el promedio")
	} else {
		for _, valor := range numeros {
			if valor < 0 {
				err = errors.New("Error: Las notas no pueden ser negativas...")
				promedio = float32(0)
				break
			}
			promedio += valor
		}
		promedio = promedio / float32(len(numeros))
	}
	return
}
