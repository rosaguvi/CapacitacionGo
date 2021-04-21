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
	fmt.Println ("Despues de Eliminar " )
	
	for key , value := range monedas {
		fmt.Println ("clave: " , key , " Valor: ", value )
	}
	
	arreglo := []int{1,2,3,4,5,6,7,8,9,10}
	
	for _ , el := range arreglo {	
		
		fmt.Println (" Elemento: ", el)
	}
}