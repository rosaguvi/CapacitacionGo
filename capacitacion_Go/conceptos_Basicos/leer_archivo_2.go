package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Inicio del Main..")
	otraFuncion()
	fmt.Println("Finalizando el Main..")
}

func otraFuncion() {
	fmt.Println("Inicio de la Otra Función..")
	leerArchivo()
	fmt.Println("Finalizando la otra Función..")

}

func leerArchivo() {
	archivo, err := os.Open("./holaa.txt")
	defer func() {
		fmt.Println("\n Cerrando el Archivo..")
		archivo.Close()
		if r := recover(); r != nil {
			fmt.Println("Recuperando el Flujo ", r)
		}
	}()
	if err != nil {
		fmt.Println("Algo salio Mal...")
		panic(err)
	}
	scanner := bufio.NewScanner(archivo)
	for i := 1; scanner.Scan(); {
		linea := scanner.Text()
		fmt.Printf("Linea %d -> %s", i, linea)
		i++
		panic(errors.New("Simulando Error"))
	}
	archivo.Close()
}
