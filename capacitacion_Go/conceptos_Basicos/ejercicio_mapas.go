package main 

import (
	"fmt"
	"strings"
)

func  main()  {
	map_palabras := make(map[string]int)
	cadena := "Esta es es una cadena de prueba para el ejercicio" 
	fmt.Println (" Por favor Ingrese una cadena de texto para evaluar: ")
	fmt.Scanf ("%s \n" , &cadena)
	map_palabras = crear_mapa(cadena);
	for palabra , cantidad := range map_palabras {
		fmt.Println("Palabra: " , palabra , "Cantidad: " , cantidad)  
	}
	pal , cant := mas_veces(map_palabras)
	fmt.Println("Palabra mas veces: " , pal , "Cantidad de Veces: " , cant )  
}

func crear_mapa(cadena string) (map[string]int) {
	palabras := strings.Fields(cadena)
	map_genpal := make(map[string]int)
	for _ , palabra := range palabras {
		_,existe:= map_genpal[palabra]
		if(existe){
			map_genpal[palabra] += 1 
		}else{
			map_genpal[palabra] = 1 
		}	  
	}
	return map_genpal 
}

func mas_veces (mapa map[string]int) (string, int) {
	veces := 0 ;
	palabra_sel := ""
	for palabra , cantidad := range mapa {
		if cantidad > veces {
			veces = cantidad 
			palabra_sel = palabra
		}		
	}
	return palabra_sel , veces
}