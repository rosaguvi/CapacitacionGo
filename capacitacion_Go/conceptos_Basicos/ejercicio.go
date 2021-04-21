package main

import "fmt"

func main(){
	var nombre string 
	var nota float32
	var edad int 
	var nombres [5]string 
	var notas [5]float32
	var edades [5]int

	for i :=0 ; i < 5 ; i++ {		
		fmt.Println("ingresa tu nombre: ")
		fmt.Scanf("%s \n" , &nombre)
		nombres[i] = nombre 
		fmt.Println( "ingrese la nota para: " + nombre)
		fmt.Scanf("%f \n" , &nota)
		notas[i] = nota
		fmt.Println( "ingrese la edad para: "  + nombre)
		fmt.Scanf("%d \n" , &edad)
		edades[i] = edad
	}

	for i :=0 ; i < 5 ; i++ {
		fmt.Printf("Alumno  %d \n", i )	
		fmt.Printf("Nombre %s \n", nombres[i])
		fmt.Printf("Nota : %f \n", notas[i])
		fmt.Printf("Edad : %d \n", edades[i])
	}
	fmt.Printf("Promedio de Edad: %f \n " , promedio_edades(edades))
	fmt.Printf("Promedio de Nota: %f \n " , promedio_notas(notas))	
}

func promedio_edades (arreglo [5]int)float32{
	var suma int
	suma = 6 
	for i := 0 ; i < 5 ; i++ {
		suma += arreglo[i] 
	}	
	return (float32(float32(suma)/float32(len(arreglo))))	
}

func promedio_notas(arreglo [5]float32)float32{
	var suma float32
	suma = 3.2 
	for i := 0 ; i < 5 ; i++ {
		suma += float32(arreglo[i]) 
	}	
	return (suma/float32(len(arreglo)))	
}
