package main

import (
	"fmt"
	"strconv"
)

type Persona struct{
	nombre, apellido string
	edad int
}

type empleado struct {
	idEmpleado string
	fecha string
}

func (e *empleado) registrarIngreso(){
	e.fecha = "21/03/2020"
}

func (p *Persona) toString () string {
	return "Nombre: " + p.nombre + " Apellido: " + p.apellido + ", Edad: " + strconv.Itoa(p.edad)
}

type Usuario struct {
	Persona
	empleado
	email string
	password string
}

func (u *Usuario) imprimirDatos() string {
	return  "Nombre: " + u.nombre + " Apellido: " + u.apellido + ", Fecha: " + u.fecha
}

func (u *Usuario) iniciarSesion() bool{
	// validar email y password
	return true 
}
func main(){
 	usuario := Usuario{Persona{nombre: "Rosmira"}, empleado{"id_001" , "10/10/2021"}, "rosmira@mail.com", "cclave"}
	fmt.Println(usuario)
	fmt.Println(usuario.nombre)
	fmt.Println(usuario.idEmpleado)
	fmt.Println(usuario.idEmpleado)
	usuario.registrarIngreso()
	fmt.Println(usuario.toString())
	fmt.Println(usuario.imprimirDatos())	
}