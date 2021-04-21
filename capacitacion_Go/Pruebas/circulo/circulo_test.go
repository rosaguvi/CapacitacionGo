package main

import "testing"

func TestCalcularPerimetro(t *testing.T) {
	t.Run("Probar cuando el valor es correcto", func(t *testing.T) {
		resultado, _ := calcularPerimetro(6.0)
		esperado := float32(37.699112)
		if resultado != esperado {
			t.Errorf("Obtuvimos %f pero esperabamos %f", resultado, esperado)
		}
	})
	t.Run("Probar cuando el radio es negativo", func(t *testing.T) {
		_, e := calcularPerimetro(-6.0)
		if e == nil {
			t.Errorf("Se esperaba error por numero negativo")
		}
	})

}
func TestCalcularArea(t *testing.T) {
	t.Run("Probar cuando el valor es correcto", func(t *testing.T) {
		resultado, _ := calcularArea(6.0)
		esperado := float32(113.097336)
		if resultado != esperado {
			t.Errorf("Obtuvimos %f pero esperabamos %f", resultado, esperado)
		}
	})
	t.Run("Probar cuando el radio es negativo", func(t *testing.T) {
		_, e := calcularArea(-6.0)
		if e == nil {
			t.Errorf("Se esperaba error por numero negativo")
		}
	})

}
