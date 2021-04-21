package main

import (
	"fmt"
	"strconv"
)

func main(){
	var nombre string 
	var x int 
	x= 2
	fmt.Println(x)
	var bandera bool
	fmt.Println (bandera)
    var strx string
	strx = strconv.Itoa(x)
	fmt.Println ("el valor de X es:" + strx)
	fmt.Printf("El valor de X es %d" , x )
	// %2f imprime decimales con 2 decimales 
	// %t para los bool
	// %S para los string
	// %d  para enteros
}

// https://play.golang.org/