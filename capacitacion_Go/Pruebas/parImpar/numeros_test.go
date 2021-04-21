package main

import "testing"

func Test_validarNumero(t *testing.T) {
	t.Run("validacion correcto con numero impar ", func(t *testing.T) {
		resultado, _ := validarNumeros(101)
		esperado := 1
		if esperado != resultado {
			t.Errorf(" Se esperaba %d y se recibio %d", esperado, resultado)
		}
	})
	t.Run("validacion correcto con numero par ", func(t *testing.T) {
		resultado, _ := validarNumeros(50)
		esperado := 0
		if esperado != resultado {
			t.Errorf(" Se esperaba %d y se recibio %d", esperado, resultado)
		}
	})
	t.Run("validacion correcto con 0 ", func(t *testing.T) {
		resultado, _ := validarNumeros(0)
		esperado := 0
		if esperado != resultado {
			t.Errorf(" Se esperaba %d y se recibio %d", esperado, resultado)
		}
	})
	t.Run("validacion Error Negativo ", func(t *testing.T) {
		_, e := validarNumeros(-20)
		if e == nil {
			t.Errorf(" Se esperaba error de numero negativo")
		}
	})
}
