package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"

	. "github.com/onsi/ginkgo" // BDD
	. "github.com/onsi/gomega" // Asserts

	// paquetes Externos
	"adminModPerl/handlers"

	"github.com/gorilla/mux"
)

func Router(ruta, metodo string, funcion http.HandlerFunc) *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc(ruta, funcion).Methods(metodo)
	return router
}

var _ = Describe("crear un usuario en la base de datos", func() {
	// se loguea el usuario en el sistema
	var token string
	BeforeEach(func() {
		var usuario = []byte(`{"email":"Pablo@personalsoft.com" , "password" : "123456" }`)
		request, _ := http.NewRequest("GET", "/", bytes.NewBuffer(usuario))
		response := httptest.NewRecorder()
		Router("/", "GET", handlers.Login).ServeHTTP(response, request)
		datosMapa := make(map[string]string)
		// Y ahora decodificamos pasando el apuntador
		_ = json.Unmarshal(response.Body.Bytes(), &datosMapa)
		token = datosMapa["token"]
	})
	Context("Buscar sin ingresar Criterio", func() {
		It("El codigo de Repuesta debe ser 200", func() {
			request, _ := http.NewRequest("GET", "/modulos", nil)
			request.Header.Add("Authorization", "Bearer "+token)
			response := httptest.NewRecorder()
			Router("/modulos", "GET", handlers.ListarModulos).ServeHTTP(response, request)
			fmt.Println(response)
			Ω(response.Code).Should(Equal(200))
		})
	})
})
