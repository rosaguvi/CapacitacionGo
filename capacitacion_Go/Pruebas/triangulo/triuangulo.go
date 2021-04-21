package main

import (
	"errors"
	"math"
)

func main() {

}

func validarTriangulo(lados []float32) (tipoMedidas, tipoAngulos string, err error) {

	if len(lados) != 3 {
		err = errors.New("Error: de cantidad de lados..")
		return
	}
	if (lados[0]+lados[1]) <= lados[2] || (lados[0]+lados[2]) <= lados[1] || (lados[1]+lados[2]) <= lados[0] {
		tipoMedidas = "No es un Triangulo"
		return
	}
	tipoMedidas = "escaleno"
	tipoAngulos = validarAngulos(lados)
	if lados[0] == lados[1] && lados[1] == lados[2] {
		tipoMedidas = "equilatero"
		return
	}
	if lados[0] == lados[1] || lados[0] == lados[2] || lados[1] == lados[2] {
		tipoMedidas = "isósceles"
		return
	}
	return
}

func validarAngulos(lados []float32) (tipoTriangulo string) {
	a := float64(lados[0])
	b := float64(lados[1])
	c := float64(lados[2])
	a_cuadrado := math.Pow(a, float64(2))
	b_cuadrado := math.Pow(b, float64(2))
	c_cuadrado := math.Pow(c, float64(2))
	anguloA := (math.Acos((b_cuadrado + c_cuadrado - a_cuadrado) / (2 * b * c))) * 180 / math.Pi
	anguloB := math.Acos((a_cuadrado+c_cuadrado-b_cuadrado)/(2*a*c)) * 180 / math.Pi
	anguloC := 180 - anguloA - anguloB
	// fmt.Println(" A", a)
	// fmt.Println(" B", b)
	// fmt.Println(" C", c)
	// fmt.Println(" A2", a_cuadrado)
	// fmt.Println(" B2", b_cuadrado)
	// fmt.Println(" C2", c_cuadrado)
	// fmt.Println("Angulo A", anguloA)
	// fmt.Println("Angulo B", anguloB)
	// fmt.Println("Angulo C", anguloC)
	tipoTriangulo = "rectángulo"
	if anguloA < 90 && anguloB < 90 && anguloC < 90 {
		tipoTriangulo = "acutángulo"
		return
	}
	if anguloA > 90 || anguloB > 90 || anguloC > 90 {
		tipoTriangulo = "obtusángulo"
		return
	}
	return
}
