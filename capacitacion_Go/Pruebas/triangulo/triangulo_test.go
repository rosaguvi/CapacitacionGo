package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidarTriangulo(t *testing.T) {
	const NO_TRIANGULO = "No es un Triangulo"
	const EQUILATERO = "equilatero"
	const ISOSCELES = "isósceles"
	const ESCALENO = "escaleno"
	const ACUTANGULO = "acutángulo"
	const ABTUSANGULO = "obtusángulo"
	const RECTANGULO = "rectángulo"
	// t.Run("Validar si es un triangulo exitosa", func(t *testing.T) {
	// 	lados := []float32{12.3, 10, 20}
	// 	resultado, _ := validarTriangulo(lados)
	// 	assert.Equal(t, TRIANGULO, resultado)
	// })
	t.Run("Validar si es un triangulo no exitosa", func(t *testing.T) {
		lados := []float32{0, 10, 10}
		resultado, _, _ := validarTriangulo(lados)
		assert.Equal(t, NO_TRIANGULO, resultado)
	})
	t.Run("Validar si es un triangulo equilatero", func(t *testing.T) {
		lados := []float32{10, 10, 10}
		resultado, _, _ := validarTriangulo(lados)
		assert.Equal(t, EQUILATERO, resultado)
	})
	t.Run("Validar si es un triangulo Isósceles", func(t *testing.T) {
		lados := []float32{10, 10, 15}
		resultado, _, _ := validarTriangulo(lados)
		assert.Equal(t, ISOSCELES, resultado)
	})
	t.Run("Validar si es un triangulo Escaleno", func(t *testing.T) {
		lados := []float32{5, 10, 14}
		resultado, _, _ := validarTriangulo(lados)
		assert.Equal(t, ESCALENO, resultado)
	})
	t.Run("Validar si es un triangulo Acutangulo", func(t *testing.T) {
		lados := []float32{15, 22, 17}
		_, resultado, _ := validarTriangulo(lados)
		assert.Equal(t, ACUTANGULO, resultado)
	})
	t.Run("Validar si es un triangulo obtusángulo", func(t *testing.T) {
		lados := []float32{3, 7, 8}
		_, resultado, _ := validarTriangulo(lados)
		assert.Equal(t, ABTUSANGULO, resultado)
	})
	t.Run("Validar si es un triangulo Rectangulo", func(t *testing.T) {
		lados := []float32{4, 3, 5}
		_, resultado, _ := validarTriangulo(lados)
		assert.Equal(t, RECTANGULO, resultado)
	})
	t.Run("Validar Error de cantidad de lados", func(t *testing.T) {
		lados := []float32{3, 5}
		_, _, e := validarTriangulo(lados)
		if e == nil {
			t.Errorf("Se esperaba error de cantidad de lados")
		}
	})
}
