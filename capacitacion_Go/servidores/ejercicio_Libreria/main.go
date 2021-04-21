package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	// paquetes Externos
	"github.com/gorilla/mux"

	"modCapacitacion/servidores/ejercicio_Libreria/database"
	"modCapacitacion/servidores/ejercicio_Libreria/models"
	"modCapacitacion/servidores/ejercicio_Libreria/respuestas"
)

const PORT = ":5000"

func main() {
	database.Migrar()
	router := mux.NewRouter()
	log.Println("Escuchando en el puerto", PORT)
	router.HandleFunc("/libros", listarLibros).Methods("GET")
	router.HandleFunc("/libro/{libro}", verLibro).Methods("GET")
	router.HandleFunc("/libro", crearLibro).Methods("POST")
	router.HandleFunc("/libro/{libro}", actualizarLibro).Methods("PUT")
	router.HandleFunc("/consultarLibrosAutor", consultarLibrosAutor).Methods("GET")
	router.HandleFunc("/consultarLibrosEditorial", consultarLibrosEditorial).Methods("GET")
	router.HandleFunc("/consultarLibrosAñopublicacion", consultarLibrosApublicacion).Methods("GET")
	router.HandleFunc("/consultarLibrosCopiasVendidas", consultarLibrosCopVendidas).Methods("GET")
	router.HandleFunc("/libro/{libro}", eliminarLibro).Methods("DELETE")
	http.Handle("/", router)
	log.Fatal(http.ListenAndServe(PORT, nil))
}

func listarLibros(w http.ResponseWriter, r *http.Request) {
	libros := database.GetLibros()
	respuestas.ResponderJSON(libros, w)
}
func consultarLibrosAutor(w http.ResponseWriter, r *http.Request) {
	var body models.DatConsulta
	var libros []models.Libro
	err := json.NewDecoder(r.Body).Decode(&body)
	validarError(err, "Error convirtiendo body....")
	libros = database.ConsultarLibroXAutor(body.Parametro)
	respuestas.ResponderJSON(libros, w)
}
func consultarLibrosEditorial(w http.ResponseWriter, r *http.Request) {
	var body models.DatConsulta
	var libros []models.Libro
	err := json.NewDecoder(r.Body).Decode(&body)
	validarError(err, "Error convirtiendo body....")
	log.Print(body.Parametro)
	libros = database.ConsultarLibroXEditorial(body.Parametro)
	respuestas.ResponderJSON(libros, w)
}
func consultarLibrosApublicacion(w http.ResponseWriter, r *http.Request) {
	var body models.DatConsulta
	var libros []models.Libro
	err := json.NewDecoder(r.Body).Decode(&body)
	validarError(err, "Error convirtiendo body....")
	log.Print(body)
	if body.Rango_inicial > 0 && body.Rango_final > 0 {
		libros = database.ConLibrosAPrango(body.Rango_inicial, body.Rango_final)
	} else {
		aho, err := strconv.Atoi(body.Parametro)
		if responderSiHayError(err, "El parametro no es un numero Valido", -1, w) {
			return
		}
		libros = database.ConLibrosApublicacion(aho)
	}
	respuestas.ResponderJSON(libros, w)
}
func consultarLibrosCopVendidas(w http.ResponseWriter, r *http.Request) {
	var body models.DatConsulta
	cntCopias, err := strconv.Atoi(body.Parametro)
	log.Print(body)
	if responderSiHayError(err, "El parametro no es un numero Valido", -1, w) {
		return
	}
	libros := database.ConLibrosNumCopVendidas(cntCopias)
	respuestas.ResponderJSON(libros, w)
}
func verLibro(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idLibro, err := strconv.Atoi(vars["libro"])
	if responderSiHayError(err, "Ide de Libro no Valido", -1, w) {
		return
	}
	libro := database.VerLibro(uint(idLibro))
	respuestas.ResponderJSON(libro, w)
}

func crearLibro(w http.ResponseWriter, r *http.Request) {
	var body models.Libro
	err := json.NewDecoder(r.Body).Decode(&body)
	validarError(err, "Error convirtiendo body....")
	log.Print(body)
	libro := database.CrearLibro(body)
	respuestas.ResponderJSON(libro, w)
}
func actualizarLibro(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idLibro, err := strconv.Atoi(vars["libro"])
	if responderSiHayError(err, "Ide de Libro no Valido", -1, w) {
		return
	}
	var body models.Libro
	err = json.NewDecoder(r.Body).Decode(&body)
	validarError(err, "Error convirtiendo body....")
	log.Print(body)
	//io.WriteString(w, " Menu de Hoteles \n")
	libro := database.EditarLibro(uint(idLibro), body)
	respuestas.ResponderJSON(libro, w)
}
func eliminarLibro(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idLibro, err := strconv.Atoi(vars["libro"])
	if responderSiHayError(err, "ID no Valido", -1, w) {
		return
	}
	database.EliminarLibro(uint(idLibro))
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
