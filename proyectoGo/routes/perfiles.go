package routes

import (
	. "adminModPerl/handlers"
	. "adminModPerl/middlewares"
)

func registrarRutasPerfil() {
	router.HandleFunc("/perfiles/{id}", Autenticar(ValidarPerfil(ListarPerfiles))).Methods("GET")         // Lista Uno o Todos
	router.HandleFunc("/perfiles", Autenticar(ValidarPerfil(ListarPerfiles))).Methods("GET")              // Lista Uno o Todos
	router.HandleFunc("/perfiles", Autenticar(ValidarPerfil(CrearPerfil))).Methods("POST")                // Crea
	router.HandleFunc("/perfiles/{id}", Autenticar(ValidarPerfil(ActualizarPerfil))).Methods("PUT")       // Actualiza
	router.HandleFunc("/perfiles/{id}", Autenticar(ValidarPerfil(EliminarPerfil))).Methods("DELETE")      // Elimina
	router.HandleFunc("/perfiles/{id}", Autenticar(ValidarPerfil(EliminarModuloPerfil))).Methods("PATCH") // Quitar Modulos al Perfil
}
