package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"

	. "github.com/onsi/ginkgo" // BDD
	. "github.com/onsi/gomega" // Asserts

	// paquetes Externos
	"github.com/gorilla/mux"
)

func Router(ruta, metodo string, funcion http.HandlerFunc) *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc(ruta, funcion).Methods(metodo)
	return router
}

var _ = Describe("crear un usuario en la base de datos", func() {

	Context("Buscar sin ingresar Criterio", func() {
		It("El codigo de Repuesta debe ser 200", func() {
			request, _ := http.NewRequest("GET", "/", nil)
			response := httptest.NewRecorder()
			Router("/", "GET", RootEndpoint).ServeHTTP(response, request)
			fmt.Println(response)
			Ω(response.Code).Should(Equal(200))
		})
	})
})
