package routes

import (
	. "adminModPerl/handlers"
	. "adminModPerl/middlewares"
)

func registrarRutasUsuario() {
	router.HandleFunc("/usuarios/{id}", Autenticar(ValidarPerfil(ListarUsuarios))).Methods("GET")     // Get All or search
	router.HandleFunc("/usuarios", Autenticar(ValidarPerfil(ListarUsuarios))).Methods("GET")          // Get All or search
	router.HandleFunc("/usuarios", Autenticar(ValidarPerfil(CrearUsuario))).Methods("POST")           // Create
	router.HandleFunc("/usuarios/{id}", Autenticar(ValidarPerfil(ActualizarUsuario))).Methods("PUT")  // Actualizar
	router.HandleFunc("/usuarios/{id}", Autenticar(ValidarPerfil(EliminarUsuario))).Methods("DELETE") // Eliminar
	router.HandleFunc("/usuario", Autenticar(ConsultarModulosUsuario)).Methods("GET")
}
