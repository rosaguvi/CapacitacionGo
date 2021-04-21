package routes

import (
	. "adminModPerl/handlers"
	. "adminModPerl/middlewares"
)

func registrarRutasPerfil() {
	router.HandleFunc("/perfiles/{id}", Autenticar(ValidarPerfil(ListarPerfiles))).Methods("GET")         // Get All or search
	router.HandleFunc("/perfiles", Autenticar(ValidarPerfil(ListarPerfiles))).Methods("GET")              // Get All or search
	router.HandleFunc("/perfiles", Autenticar(ValidarPerfil(CrearPerfil))).Methods("POST")                // Create
	router.HandleFunc("/perfiles/{id}", Autenticar(ValidarPerfil(ActualizarPerfil))).Methods("PUT")       // Actualizar
	router.HandleFunc("/perfiles/{id}", Autenticar(ValidarPerfil(EliminarPerfil))).Methods("DELETE")      // Eliminar
	router.HandleFunc("/perfiles/{id}", Autenticar(ValidarPerfil(EliminarModuloPerfil))).Methods("PATCH") // Quitar Modulos al Perfil
}
