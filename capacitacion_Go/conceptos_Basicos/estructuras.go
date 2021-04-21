package main

import (
	"fmt"
	"strconv"
)

type Humano struct{
	nombre string
	edad int
	peso int
	sexo string
}
type Grupo struct{
	integrantes []Humano
}

func (humano Humano) imprimir() string{
	return "Nombre: " + humano.nombre + ", Edad: " + strconv.Itoa(humano.edad)
}

func (humano *Humano) cambiarValor(nombre string, edad int){
	humano.nombre = nombre 
	humano.edad = edad
}

func main(){
	persona := Humano {nombre: "rosmira"}
	fmt.Println (persona.imprimir())
	persona2:= new(Humano)
	fmt.Println (persona2.imprimir())
	persona2.nombre = "juan"
	persona2.edad = 30
	persona2.peso = 60 
	fmt.Println (persona2)
	persona.cambiarValor("Julian", 22 )
	fmt.Println (persona.imprimir())
	persona2.cambiarValor("Maria", 30 )
	fmt.Println (persona2.imprimir())
	// grupo := mew(Grupo)
	// fmt.Println (grupo)
}