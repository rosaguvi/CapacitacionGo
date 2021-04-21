package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidarNumero(t *testing.T) {
	const PALINDROMO = "palíndromo"
	const NO_PALINDROMO = "no palíndromo"
	t.Run("validar palindromo", func(t *testing.T) {
		resultado, _ := validarNumero(12321)
		assert.Equal(t, resultado, PALINDROMO)
	})
	t.Run("validar palindromo no exitoso", func(t *testing.T) {
		resultado, _ := validarNumero(12324)
		assert.Equal(t, resultado, NO_PALINDROMO)
	})
	t.Run("validar error de numero negativo", func(t *testing.T) {
		_, e := validarNumero(-12324)
		if e == nil {
			t.Errorf("se error de numero negativo")
		}
	})
}
