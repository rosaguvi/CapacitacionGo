package main

import "testing"

func Test_CalHipotenusa(t *testing.T) {
	t.Run("calculo correcto con lados enteros", func(t *testing.T) {
		resultado, _ := calcularHipotenusa(5, 6)
		esperado := float32(7.81024967590665)
		if resultado != esperado {
			t.Errorf("se esperaba %f y se recibio %f", esperado, resultado)
		}
	})
	t.Run("calculo correcto con lados flotantes", func(t *testing.T) {
		resultado, _ := calcularHipotenusa(5.32, 6.25)
		esperado := float32(8.20761232027927)
		if resultado != esperado {
			t.Errorf("se esperaba %f y se recibio %f", esperado, resultado)
		}
	})
	t.Run("calculo correcto con lados flotantes pequeños", func(t *testing.T) {
		resultado, _ := calcularHipotenusa(0.32, 0.25)
		esperado := float32(0.406078810084939)
		if resultado != esperado {
			t.Errorf("se esperaba %f y se recibio %f", esperado, resultado)
		}
	})
	t.Run("Error por lados negativos", func(t *testing.T) {
		_, e := calcularHipotenusa(-5.32, 6.25)
		if e == nil {
			t.Errorf("se esperaba error de lados negativos")
		}
	})
	t.Run("Error por lados cero", func(t *testing.T) {
		_, e := calcularHipotenusa(0, 6.25)
		if e == nil {
			t.Errorf("se esperaba error de lados negativos o cero")
		}
	})
}
