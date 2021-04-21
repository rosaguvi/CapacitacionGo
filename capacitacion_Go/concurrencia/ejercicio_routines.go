// # Generar un numero aleatorio n de 10 a 30
// # generar n gorutines
// # por cada gorutine, hacer un sleep de msegundo, donde m es un aleatorio de 1 - 5
// # dentro de la gorutine imprimir "soy la rutina {n}"

package main

import (
	"fmt"
	"math/rand"
	"time"
)

var src = rand.NewSource(time.Now().UnixNano())
var r = rand.New(src)

func main() {
	fmt.Println("Iniciando el main..")
	n := 0
	for n == 0 {
		n = r.Intn(30)
	}
	fmt.Println("numero de Rutinas  ", n)
	for i := 0; i < n; i++ {
		fmt.Println("Estoy en el for  ", i)
		go func(num int) {
			m := 0
			for m == 0 {
				m = r.Intn(5)
			}
			fmt.Println("soy la rutina  ", num)
			fmt.Println("Espera  ", m, "segundos")
			time.Sleep(time.Duration(m) * time.Second)
		}(i)
	}
	fmt.Println("Finalizando el main..")
	var espera string
	fmt.Scanln(&espera)
}
