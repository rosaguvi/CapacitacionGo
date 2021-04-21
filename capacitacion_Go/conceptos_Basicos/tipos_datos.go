package main

import "fmt"

// Tipos de datos Personalizados
type minuto int
type hora int

// metodo 
func (m minuto) aEntero() int {
	return int(m) * 100 
}

func main(){
	var tiempo minuto = 3 
	fmt.Println (tiempo)
	var convertido int 
	convertido = tiempo.aEntero()
	fmt.Println (convertido)
}