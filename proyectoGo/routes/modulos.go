package routes

import (
	. "adminModPerl/handlers"
	. "adminModPerl/middlewares"
)

func registrarRutasModulo() {
	router.HandleFunc("/modulos/{id}", Autenticar(ValidarPerfil(ListarModulos))).Methods("GET")     // Lista uno o Todos
	router.HandleFunc("/modulos", Autenticar(ValidarPerfil(ListarModulos))).Methods("GET")          // Lista uno o Todos
	router.HandleFunc("/modulos", Autenticar(ValidarPerfil(CrearModulo))).Methods("POST")           // Crear
	router.HandleFunc("/modulos/{id}", Autenticar(ValidarPerfil(ActualizarModulo))).Methods("PUT")  // Actualizar
	router.HandleFunc("/modulos/{id}", Autenticar(ValidarPerfil(EliminarModulo))).Methods("DELETE") // Eliminar
}
