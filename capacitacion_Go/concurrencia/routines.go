package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	fmt.Println("Imprimiendo Nombre")
	go imprimirNombre("Rosmira")
	fmt.Println("Fin de la Impresión")
	var espera string
	fmt.Scanln(&espera)
}

func imprimirNombre(nombre string) {
	letras := strings.Split(nombre, "")
	for _, letra := range letras {
		time.Sleep(2 * time.Second)
		fmt.Println(letra)
	}
}
