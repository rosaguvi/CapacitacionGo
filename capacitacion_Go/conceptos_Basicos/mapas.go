package main

import "fmt" 

func main (){
	monedas := make(map[string]string)
	monedas ["colombia"] = "peso colombiano" 
	monedas ["eeuu"] = "Dolar" 
	monedas ["españa"] = "Euro" 
	fmt.Println (monedas)
	fmt.Println (len(monedas))
	for key , value := range monedas {
		fmt.Println ("clave: " , key , " Valor: ", value )
	}
	delete(monedas , "españa")

	for key , value := range monedas {
		fmt.Println ("clave: " , key , " Valor: ", value )
	}
}