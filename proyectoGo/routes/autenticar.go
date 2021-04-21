package routes

import (
	. "adminModPerl/handlers"
)

func registrarRutasAutenticar() {
	router.HandleFunc("/autenticar", Login).Methods("POST")
}
