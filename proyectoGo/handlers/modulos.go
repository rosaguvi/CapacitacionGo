package handlers

import (
	"net/http"

	"github.com/gorilla/mux"

	"adminModPerl/io/request"
	"adminModPerl/io/response"
	"adminModPerl/models"
	"adminModPerl/services"
)

func ListarModulos(w http.ResponseWriter, r *http.Request) {
	status := http.StatusOK
	vars := mux.Vars(r)
	err, permisos := ConsultarUsuarioPerfil(r)
	if err != nil {
		response.Error(err.Error(), http.StatusBadRequest, w)
		return
	}
	listModulos, e := services.ListarModulos(vars, permisos.ID, permisos.Rol)
	if e != nil {
		response.Error(e.Error(), http.StatusBadRequest, w)
		return
	}
	if len(listModulos) <= 0 || listModulos[0].ID == 0 {
		status = http.StatusNoContent
	}
	response.Json(listModulos, status, w)
}
func CrearModulo(w http.ResponseWriter, r *http.Request) {
	var body models.Modulo
	err, permisos := ConsultarUsuarioPerfil(r)
	if err != nil {
		response.Error(err.Error(), http.StatusBadRequest, w)
		return
	}
	err = request.Json(r, &body)
	if err != nil {
		response.Error("El Formato del Body no es el Correcto", http.StatusBadRequest, w)
		return
	}
	body.UsuarioGB = permisos.ID
	e, modulo := services.CrearModulo(body, permisos.Rol)
	if e != nil {
		response.Error(e.Error(), http.StatusBadRequest, w)
		return
	}
	response.Json(modulo, http.StatusOK, w)
}

func ActualizarModulo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var body models.Modulo
	err, permisos := ConsultarUsuarioPerfil(r)
	if err != nil {
		response.Error(err.Error(), http.StatusBadRequest, w)
		return
	}
	err = request.Json(r, &body)
	if err != nil {
		response.Error("El Formato del Body no es el Correcto", http.StatusBadRequest, w)
		return
	}
	e, Modulo := services.ActualizarModulo(body, vars, permisos.ID, permisos.Rol)
	if e != nil {
		response.Error(e.Error(), http.StatusBadRequest, w)
		return
	}
	response.Json(Modulo, http.StatusOK, w)
}
func EliminarModulo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	err, permisos := ConsultarUsuarioPerfil(r)
	if err != nil {
		response.Error(err.Error(), http.StatusBadRequest, w)
		return
	}
	e := services.EliminarModulo(vars, permisos.ID, permisos.Rol)
	if e != nil {
		response.Error(e.Error(), http.StatusBadRequest, w)
		return
	}
	response.Json("Modulo Eliminado Correctamente", http.StatusOK, w)
}
