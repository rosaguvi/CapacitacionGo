package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Curso struct {
	Id          int
	Nombre      string
	Descripcion string
}

const PUERTO = ":5000"

func main() {
	log.Println("Corriendo en el puerto", PUERTO)

	http.HandleFunc("/redirect", redireccion)
	http.HandleFunc("/hola", hola)
	http.HandleFunc("/not-found", noEncontrado)
	http.HandleFunc("/error", errorPeticion)
	http.HandleFunc("/parametros", parametros)
	http.HandleFunc("/enviar-json", enviarJson)

	log.Fatal(http.ListenAndServe(PUERTO, nil))
}

func hola(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hola Mundo!")
	fmt.Fprint(w, "Hola con FMT")
}
func redireccion(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Redirigiendo a Hola")
	http.Redirect(w, r, "/hola", 301)
}

func noEncontrado(w http.ResponseWriter, r *http.Request) {
	http.NotFound(w, r)
}

func errorPeticion(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Error en el servidor", 502)
}

func parametros(w http.ResponseWriter, r *http.Request) {
	// log.Println(r.URL)
	// log.Println(r.URL.RawQuery)
	mapaParametros := r.URL.Query()
	log.Println(mapaParametros)
	// log.Println(mapaParametros["idProducto"][0])
	fmt.Fprintln(w, "Recibiendo parámetros")
}

func enviarJson(w http.ResponseWriter, r *http.Request) {
	curso1 := Curso{1, "Curso de Go", "Curso de introducción a Go (Golang)"}
	curso2 := Curso{2, "Curso de Flutter", "Curso de introducción a Flutter"}
	curso3 := Curso{3, "Curso de Angular", "Curso de introducción a Angular"}
	cursos := []Curso{curso1, curso2, curso3}
	json.NewEncoder(w).Encode(cursos)
}
