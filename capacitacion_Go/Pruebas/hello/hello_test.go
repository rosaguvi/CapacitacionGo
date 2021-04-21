package main

import "testing"

func TestSaludar(t *testing.T) {
	t.Run("Probar cuando el nombre está vacío", func(t *testing.T) {
		saludo := saludar("")
		esperado := "Hola Mundo!"
		if saludo != esperado {
			t.Errorf("Obtuvimos %s pero esperabamos %s", saludo, esperado)
		}
	})
	t.Run("Probar cuando el nombre no esté vacío", func(t *testing.T) {
		saludo := saludar("Juan")
		esperado := "Hola Juan"
		if saludo != esperado {
			t.Errorf("Obtuvimos %s pero esperabamos %s", saludo, esperado)
		}
	})

}

// func TestSaludar(t *testing.T) {
// 	saludo := saludar()
// 	esperado := "Hola Mundo"
// 	if saludo != esperado {
// 		t.Errorf("Obtuvimos %s pero esperabamos %s", saludo, esperado)
// 	}
// }
