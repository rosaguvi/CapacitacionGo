// # Generar un numero aleatorio n de 10 a 30
// gnerar un amatriz de n filas * 5 columnas
// # generar n gorutines para llenar cada fila

package main

import (
	"fmt"
	"math/rand"
	"time"
)

var src = rand.NewSource(time.Now().UnixNano())
var r = rand.New(src)
var matriz [][]int

func main() {
	fmt.Println("Iniciando el main..")
	n := 0
	for n == 0 {
		n = r.Intn(10)
	}
	fmt.Println("numero de Rutinas  y filas de la matriz", n)
	//var arreglo [5]int
	matriz = make([][]int, n)
	defer func() {
		for fila := 0; fila < n; fila++ {
			fmt.Printf("[")
			for col := 0; col < 5; col++ {
				fmt.Print(" ")
				fmt.Print(matriz[fila][col])
			}
			fmt.Printf(" ] \n")
		}
	}()
	//fmt.Println(arreglo)
	//matriz = append(matriz, arreglo)
	//matriz = append(matriz, arreglo)
	//fmt.Println(matriz)
	for i := 0; i < n; i++ {
		fmt.Println("Estoy en el for  ", i)
		go llenarMatriz(i)
	}
	// for contador < n {
	// 	time.Sleep(1 * time.Second)
	// }
	//fmt.Println(matriz)

	fmt.Println("Finalizando el main..")

}

func llenarMatriz(fila int) {
	fmt.Println("soy la rutina  ", fila, "Llenando la fila")
	matriz[fila] = make([]int, 5)
	for j := 0; j < 5; j++ {
		matriz[fila][j] = r.Intn(100)
	}

}
