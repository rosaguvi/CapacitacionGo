package main

import (
	"log"
	"net/http"

	// paquetes Externos
	"github.com/gorilla/mux"
)

func RootEndpoint(response http.ResponseWriter, request *http.Request) {
	response.WriteHeader(200)
	response.Write([]byte("Hello World"))
}

const PORT = ":5000"

func main() {
	router := mux.NewRouter()
	log.Println("Escuchando en el puerto", PORT)
	router.HandleFunc("/", RootEndpoint).Methods("GET")
	log.Fatal(http.ListenAndServe(PORT, router))
}
