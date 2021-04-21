package main

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFizzBuzz(t *testing.T) {
	const MSG_NO_VALIDO = "No es un numero valido"
	const FIZZ = "fizz"
	const BUZZ = "buzz"
	const FIZZBUZZ = "fizzbuzz"
	t.Run("Probar cuando se envia 0", func(t *testing.T) {
		salida := fizzBuzz(0)
		assert.Equal(t, salida, MSG_NO_VALIDO)
		// if salida != MSG_NO_VALIDO {
		// 	t.Errorf("No es un numero Valido")
		// }
	})
	t.Run("Probar cuando se envia negativo", func(t *testing.T) {
		salida := fizzBuzz(-5)
		assert.Equal(t, salida, MSG_NO_VALIDO)
		// if salida != MSG_NO_VALIDO {
		// 	t.Errorf("No es un numero Valido")
		// }
	})
	t.Run("Probar cuando se envia un multiplo de 3", func(t *testing.T) {
		salida := fizzBuzz(9)
		assert.Equal(t, salida, FIZZ)
		//
	})
	t.Run("Probar cuando se envia un multiplo de 5", func(t *testing.T) {
		salida := fizzBuzz(20)
		assert.Equal(t, salida, BUZZ)
		// if salida != BUZZ {
		// 	t.Errorf("Se esperaba %s y se recibio %s", BUZZ, salida)
		// }
	})
	t.Run("Probar cuando se envia un multiplo de 3 y 5", func(t *testing.T) {
		salida := fizzBuzz(15)
		assert.Equal(t, salida, FIZZBUZZ)
		// if salida != FIZZBUZZ {
		// 	t.Errorf("Se esperaba %s y se recibio %s", FIZZBUZZ, salida)
		// }
	})
	t.Run("Probar cuando se envia un numero no multiplo de 3 y 5", func(t *testing.T) {
		salida := fizzBuzz(22)
		esperado := strconv.Itoa(22)
		assert.Equal(t, salida, esperado)
		// if salida != esperado {
		// 	t.Errorf("Se esperaba %s y se recibio %s", esperado, salida)
		// }
	})
}
