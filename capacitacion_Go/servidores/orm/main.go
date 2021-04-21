package main

import (
	//propios de go
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	// paquetes Externos
	"github.com/gorilla/mux"

	//paquetes desarrollador
	"modCapacitacion/servidores/orm/database"
	"modCapacitacion/servidores/orm/models"
	"modCapacitacion/servidores/orm/respuestas"
)

const PORT = ":5000"

func main() {
	database.Migrar()
	router := mux.NewRouter()
	log.Println("Escuchando en el puerto", PORT)
	router.HandleFunc("/cursos", listarCursos).Methods("GET")
	router.HandleFunc("/curso/{curso}", verCurso).Methods("GET")
	router.HandleFunc("/curso", crearCurso).Methods("POST")
	router.HandleFunc("/curso/{curso}", actualizarCurso).Methods("PUT")
	router.HandleFunc("/curso/{curso}", eliminarCurso).Methods("DELETE")
	http.Handle("/", router)
	log.Fatal(http.ListenAndServe(PORT, nil))
}

func listarCursos(w http.ResponseWriter, r *http.Request) {
	cursos := database.GetCursos()
	respuestas.ResponderJSON(cursos, w)
}
func crearCurso(w http.ResponseWriter, r *http.Request) {
	var body models.Curso
	err := json.NewDecoder(r.Body).Decode(&body)
	validarError(err, "Error convirtiendo body....")
	curso := database.CrearCurso(body.Nombre, body.Descripcion)
	respuestas.ResponderJSON(curso, w)
}
func verCurso(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idCurso, err := strconv.Atoi(vars["curso"])
	if responderSiHayError(err, "Ide no Valido", -1, w) {
		return
	}
	curso := database.VerCurso(uint(idCurso))
	respuestas.ResponderJSON(curso, w)
}
func actualizarCurso(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idCurso, err := strconv.Atoi(vars["curso"])
	if responderSiHayError(err, "ID no Valido", -1, w) {
		return
	}
	var body models.Curso
	err = json.NewDecoder(r.Body).Decode(&body)
	validarError(err, "Error convirtiendo body....")
	curso := database.EditarCurso(uint(idCurso), body.Nombre, body.Descripcion)
	respuestas.ResponderJSON(curso, w)
}
func eliminarCurso(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idCurso, err := strconv.Atoi(vars["curso"])
	if responderSiHayError(err, "ID no Valido", -1, w) {
		return
	}
	database.EliminarCurso(uint(idCurso))
	respuestas.ResponderJSON(models.Respuesta{Mensaje: "Ok", Codigo: 1}, w)
}

func responderSiHayError(err error, mensaje string, codigo int, w http.ResponseWriter) bool {
	if err != nil {
		respuestas.ResponderJSON(
			models.Respuesta{Mensaje: mensaje, Codigo: codigo},
			w)
		return true
	}
	return false
}
func validarError(err error, mensaje string) {
	if err != nil {
		log.Print(mensaje)
		panic(err)
	}
}
