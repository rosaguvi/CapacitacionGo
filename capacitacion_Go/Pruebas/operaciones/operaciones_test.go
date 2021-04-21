package main

import "testing"

func TestSuma(t *testing.T) {
	t.Run("Probar que el resultado es correcto", func(t *testing.T) {
		resultado, _ := sumar(3, 4)
		esperado := 7
		if resultado != esperado {
			t.Errorf("Obtuvimos %d pero esperabamos %d", resultado, esperado)
		}
	})
	t.Run("Probar que los resultados son diferentes", func(t *testing.T) {
		resultado, _ := sumar(3, 4)
		esperado := 8
		if resultado == esperado {
			t.Errorf("la suma esta mal")
		}
	})
	t.Run("Probar error cuando X o Y son menores a cero", func(t *testing.T) {
		_, e := sumar(-1, 1)
		if e == nil {
			t.Errorf("se esperaba error por numeros menores a cero")
		}
	})
}
func TestResta(t *testing.T) {
	t.Run("Probar que el resultado es correcto", func(t *testing.T) {
		resultado, _ := restar(10, 5)
		esperado := 5
		if resultado != esperado {
			t.Errorf("Obtuvimos %d pero esperabamos %d", resultado, esperado)
		}
	})
	t.Run("Probar que los resultados son diferentes", func(t *testing.T) {
		resultado, _ := sumar(8, 3)
		esperado := 4
		if resultado == esperado {
			t.Errorf("se esperaba resultado diferente")
		}
	})
	t.Run("Probar error cuando X o Y son menores a cero", func(t *testing.T) {
		_, e := restar(-1, 1)
		if e == nil {
			t.Errorf("se esperaba error por numeros menores a cero")
		}
	})
	t.Run("Probar Y mayor que X", func(t *testing.T) {
		_, e := restar(5, 10)
		if e == nil {
			t.Errorf("se esperaba error por Y mayor a X ")
		}
	})
}
func TestMultiplicar(t *testing.T) {
	t.Run("Probar que el resultado es correcto", func(t *testing.T) {
		resultado, _ := multiplicar(10, 5)
		esperado := 50
		if resultado != esperado {
			t.Errorf("Obtuvimos %d pero esperabamos %d", resultado, esperado)
		}
	})
	t.Run("Probar error cuando X o Y son menores a cero", func(t *testing.T) {
		_, e := multiplicar(-1, 1)
		if e == nil {
			t.Errorf("se esperaba error por numeros menores a cero")
		}
	})
}
func TestDividir(t *testing.T) {
	t.Run("Probar que el resultado es correcto", func(t *testing.T) {
		resultado, _ := dividir(10, 5)
		esperado := 2
		if resultado != esperado {
			t.Errorf("Obtuvimos %d pero esperabamos %d", resultado, esperado)
		}
	})
	t.Run("Probar error cuando X o Y son menores a cero", func(t *testing.T) {
		_, e := dividir(-1, 1)
		if e == nil {
			t.Errorf("se esperaba error por numeros menores a cero")
		}
	})
	t.Run("Probar error division cero", func(t *testing.T) {
		_, e := dividir(10, 0)
		if e == nil {
			t.Errorf("se esperaba error por numeros menores a cero")
		}
	})
	t.Run("Probar error division no exacta", func(t *testing.T) {
		resultado, _ := dividir(10, 3)
		esperado := 3
		if resultado != esperado {
			t.Errorf("Obtuvimos %d pero esperabamos %d", resultado, esperado)
		}
	})
}
