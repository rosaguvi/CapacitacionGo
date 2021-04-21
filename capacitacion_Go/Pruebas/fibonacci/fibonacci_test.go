package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerarfibonacci(t *testing.T) {
	t.Run("validar numero n de Fibonacci con n", func(t *testing.T) {
		valor, _ := generarFibonacci(10)
		esperado := 55
		assert.Equal(t, valor, esperado)
	})
	t.Run("validar el primer numero de la serie fibonacci", func(t *testing.T) {
		valor, _ := generarFibonacci(0)
		esperado := 0
		assert.Equal(t, valor, esperado)
	})
	t.Run("validar el segundo Numero de la serie fibonacci", func(t *testing.T) {
		valor, _ := generarFibonacci(1)
		esperado := 1
		assert.Equal(t, valor, esperado)
	})
	t.Run("validar Error de numero negativo", func(t *testing.T) {
		_, e := generarFibonacci(-2)
		if e == nil {
			t.Errorf("se esperaba error de negativos")
		}
	})

}
