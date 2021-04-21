package main

import "errors"

func main() {

}

func sumar(x, y int) (resultado int, err error) {
	if x < 0 || y < 0 {
		err = errors.New("No se pueden operar numeros menores a cero")
	} else {
		resultado = x + y
	}
	return
}
func restar(x, y int) (resultado int, err error) {
	if x < 0 || y < 0 {
		err = errors.New("No se pueden operar numeros menores a cero")
	} else if y > x {
		err = errors.New(" Y no puede ser mayor a X ")
	} else {
		resultado = x - y
	}
	return
}
func multiplicar(x, y int) (resultado int, err error) {
	if x < 0 || y < 0 {
		err = errors.New("No se pueden operar numeros menores a cero")
	} else {
		resultado = x * y
	}
	return
}

func dividir(x, y int) (resultado int, err error) {
	if x < 0 || y < 0 {
		err = errors.New("No se pueden operar numeros menores a cero")
	} else if y == 0 {
		err = errors.New(" Error: Division por cero")
	} else {
		resultado = x / y
	}
	return
}
