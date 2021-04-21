package main

import "fmt"

func main(){
	var nombre string 
	fmt.Println("ingresa t nombre: ")
	fmt.Scanf("%s" , &nombre)
	saludar(nombre)
}

func saludar(nombre string){
	fmt.Println ("Hola " + nombre)
}