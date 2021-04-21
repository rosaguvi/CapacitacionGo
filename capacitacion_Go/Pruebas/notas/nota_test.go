package main

import "testing"

func TestCalcularNota(t *testing.T) {
	t.Run("valida nota rango 3 con valor 0 ", func(t *testing.T) {
		notas := []float32{5, 3, 4}
		resultado, _ := calcularNota(notas)
		esperado := float32(0)
		if esperado != resultado {
			t.Errorf("Se esperaba %f y se recibio %f", esperado, resultado)
		}
	})
	t.Run("valida nota rango 6 con valor 0 ", func(t *testing.T) {
		notas := []float32{7, 8, 9}
		resultado, _ := calcularNota(notas)
		esperado := float32(0)
		if esperado != resultado {
			t.Errorf("Se esperaba %f y se recibio %f", esperado, resultado)
		}
	})
	t.Run("valida nota rango 3 con valor > 0 ", func(t *testing.T) {
		notas := []float32{2.5, 2.8, 1.9}
		resultado, _ := calcularNota(notas)
		esperado := float32(4.80000000)
		if esperado != resultado {
			t.Errorf("Se esperaba %f y se recibio %f", esperado, resultado)
		}
	})
	t.Run("valida nota rango 6 con valor > 0 ", func(t *testing.T) {
		notas := []float32{5.5, 4.8, 6.9}
		resultado, _ := calcularNota(notas)
		esperado := float32(6.8)
		if esperado != resultado {
			t.Errorf("Se esperaba %f y se recibio %f", esperado, resultado)
		}
	})
	t.Run("valida error notas menor a cero ", func(t *testing.T) {
		notas := []float32{-5.3, 4.8, 6.9}
		_, e := calcularNota(notas)
		if e == nil {
			t.Errorf("Se esperaba Error de nota negativa")
		}
	})
	t.Run("valida error menos cantidad de nota ", func(t *testing.T) {
		notas := []float32{4.8, 6.9}
		_, e := calcularNota(notas)
		if e == nil {
			t.Errorf("Se esperaba Error de cantidad de notas")
		}
	})
	t.Run("valida error de no pasa materia rango 3 ", func(t *testing.T) {
		notas := []float32{2.3, 3.1, 1.2}
		_, e := calcularNota(notas)
		if e == nil {
			t.Errorf("Se esperaba Error de no pasa la materia")
		}
	})
	t.Run("valida error de no pasa materia rango 6 ", func(t *testing.T) {
		notas := []float32{6.3, 3.1, 1.2}
		_, e := calcularNota(notas)
		if e == nil {
			t.Errorf("Se esperaba Error de no pasa la materia")
		}
	})
}
