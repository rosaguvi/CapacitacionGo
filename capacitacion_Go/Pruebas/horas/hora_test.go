package main

import "testing"

func TestConsultarHora(t *testing.T) {
	t.Run("valida hora correcta sin cambiar de dia ", func(t *testing.T) {
		resultado, _ := calculaHora(4)
		esperado := 21
		if esperado != resultado {
			t.Errorf("Se esperaba %d y se recibio %d", esperado, resultado)
		}
	})
	t.Run("valida hora correcta sin cambiando de dia", func(t *testing.T) {
		resultado, _ := calculaHora(12)
		esperado := 5
		if esperado != resultado {
			t.Errorf("Se esperaba %d y se recibio %d", esperado, resultado)
		}
	})
	t.Run("valida error hora negativa", func(t *testing.T) {
		_, e := calculaHora(-5)
		if e == nil {
			t.Errorf("Se esperaba error de hora negativa")
		}
	})
}
