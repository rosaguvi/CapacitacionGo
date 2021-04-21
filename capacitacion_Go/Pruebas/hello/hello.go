package main

import "fmt"

func main() {
	fmt.Println(saludar(""))
}
func saludar(nombre string) string {
	if nombre != "" {
		return "Hola " + nombre
	} else {
		return "Hola Mundo!"
	}
}
