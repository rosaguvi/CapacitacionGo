package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

type estudiante struct {
	nombre     string
	notas      []materia
	nota_final float32
	edad       int
}

func (est *estudiante) genPromedio() {
	suma := float32(0)
	for i := 0; i < len(est.notas); i++ {
		suma += est.notas[i].nota
	}
	est.nota_final = suma / float32(len(est.notas))
}

type materia struct {
	nombre string
	nota   float32
}

func main() {
	var vr_estudiantes []estudiante
	var vr_nota materia
	cnt_est := 10
	cnt_notas := 3

	for i := 0; i < cnt_est; i++ {
		var vr_est estudiante
		fmt.Println("ingresa tu nombre: ")
		fmt.Scanf("%s \n", &vr_est.nombre)
		vr_est.nombre = "Estudiante_" + strconv.Itoa(i+1)
		fmt.Println("ingreso de notas para: " + vr_est.nombre)
		for j := 0; j < cnt_notas; j++ {
			fmt.Printf("ingreso el nombre de la materia %d : ", j+1)
			fmt.Scanf("%s \n", &vr_nota.nombre)
			vr_nota.nombre = "Materia_" + strconv.Itoa(j+1)
			fmt.Printf("ingreso la nota para la Materia %s: ", vr_nota.nombre)
			fmt.Scanf("%f \n", &vr_nota.nota)
			vr_nota.nota = (rand.Float32() * 5)
			vr_est.notas = append(vr_est.notas, vr_nota)
		}
		fmt.Println("ingrese la edad para: " + vr_est.nombre)
		fmt.Scanf("%d \n", &vr_est.edad)
		vr_est.edad = (rand.Intn(100))
		vr_estudiantes = append(vr_estudiantes, vr_est)
	}
	pro_nota := float32(0)
	pro_edad := float32(0)
	pro_nota, pro_edad, vr_estudiantes = calcularPromedio(vr_estudiantes)
	imprimeEstudiantes(vr_estudiantes)
	fmt.Println("\nPromedio de las Notas:", pro_nota)
	fmt.Println("Promedio de la Edad:", pro_edad)
}

func calcularPromedio(ests []estudiante) (float32, float32, []estudiante) {
	fmt.Println(ests)
	sum_nota := float32(0)
	sum_edad := 0
	cnt_est := float32(len(ests))
	for i := 0; i < len(ests); i++ {
		sum_edad += ests[i].edad
		ests[i].genPromedio()
		sum_nota += ests[i].nota_final
	}
	return (float32(sum_nota) / cnt_est), (float32(sum_edad) / cnt_est), ests
}

func imprimeEstudiantes(est []estudiante) {
	for i := 0; i < len(est); i++ {
		fmt.Println("Datos para el Estudiante " + strconv.Itoa(i+1) + ": ")
		fmt.Println("Nombre: " + est[i].nombre)
		fmt.Printf("Edad: %d \n", est[i].edad)
		fmt.Println("******* Notas ********* ")
		for j := 0; j < len(est[i].notas); j++ {
			fmt.Printf("%s: %f \n", est[i].notas[j].nombre, est[i].notas[j].nota)
		}
		fmt.Printf("Nota final: %f \n \n", est[i].nota_final)
	}
}
