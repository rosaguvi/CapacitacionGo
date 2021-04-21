package main

import "testing"

func Test_promedio(t *testing.T) {
	t.Run("calculos Correcto con 7 numeros enteros", func(t *testing.T) {
		numeros := []float32{1, 20, 30, 45, 50, 60, 20}
		resultado, _ := calcularPromedio(numeros)
		esperado := float32(32.285714)
		if resultado != esperado {
			t.Errorf("se esperaba %f y se recibio %f", esperado, resultado)
		}

	})
	t.Run("calculos Correcto con 7 numeros enteros, primeros en cero", func(t *testing.T) {
		numeros := []float32{0, 0, 30, 45, 50, 60, 20}
		resultado, _ := calcularPromedio(numeros)
		esperado := float32(29.285715)
		if resultado != esperado {
			t.Errorf("se esperaba %f y se recibio %f", esperado, resultado)
		}
	})
	t.Run("calculos Correcto con 7 numeros float", func(t *testing.T) {
		numeros := []float32{0.25, 0.56, 30.58, 45.35, 50.56, 60.01, 20.32}
		resultado, _ := calcularPromedio(numeros)
		esperado := float32(29.6614285714286)
		if resultado != esperado {
			t.Errorf("se esperaba %f y se recibio %f", esperado, resultado)
		}
	})
	t.Run("error cuando no hay numeros para el promedio", func(t *testing.T) {
		var numeros []float32
		_, e := calcularPromedio(numeros)
		if e == nil {
			t.Errorf("se esperaba error de numeros vacios")
		}
	})
	t.Run("error cuando con numeros negativos", func(t *testing.T) {
		numeros := []float32{1, 20, 30, -45, 50, 60, 20}
		_, e := calcularPromedio(numeros)
		if e == nil {
			t.Errorf("se esperaba error de numeros positivos y negativos")
		}
	})
}
