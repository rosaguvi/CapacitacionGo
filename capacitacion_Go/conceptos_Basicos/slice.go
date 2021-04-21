package main 

import "fmt"

func main(){
	var arreglo [5]int
	var arreglo2 [10]int
	var slice []int
    	slice = arreglo [:]
	fuente := []int{10,20,30,40}
	destino := make([]int, 2)
	// destino := make([]int, 2 , 10)
	copy (destino,fuente)
	//copy (arreglo2,arreglo)
	// slice = append (slice , 10)
	//fmt.println (len(slice))
	//fmt.println (cap(slice))
	slice2 := make ([]int , 3 )
	fmt.Println (arreglo)
	fmt.Println (arreglo2)
	fmt.Println (slice)
	fmt.Println (slice2)
	fmt.Println ("fuente: ",fuente)
	fmt.Println ("Destino: " , destino)

}

// package main 

// import "fmt"

// func main(){
// 	var arreglo [5]int
// 	var arreglo2 [10]int
// 	var slice []int

// 	copy (slice,arreglo)
// 	slice = append (slice , 10)
// 	fmt.println (len(slice))
// 	fmt.println (cap(slice))
// 	slice2 := make ([]int , 3 )
// 	fmt.println(arreglo)
// 	fmt.println(arreglo2)
// 	fmt.println(slice)
// 	fmt.println(slice)
// 	fmt.Println(arreglo)
// 	fmt.Println(arreglo2 )
// }