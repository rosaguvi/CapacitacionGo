package handlers

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"adminModPerl/io/request"
	"adminModPerl/io/response"
	"adminModPerl/models"
	"adminModPerl/services"
)

func ListarPerfiles(w http.ResponseWriter, r *http.Request) {
	status := http.StatusOK
	vars := mux.Vars(r)
	err, permisos := ConsultarUsuarioPerfil(r)
	if err != nil {
		response.Error(err.Error(), http.StatusBadRequest, w)
		return
	}
	listPerfiles, e := services.ListarPerfiles(vars, permisos.ID, permisos.Rol)
	if e != nil {
		response.Error(e.Error(), http.StatusBadRequest, w)
		return
	}
	if len(listPerfiles) <= 0 || listPerfiles[0].ID == 0 {
		status = http.StatusNoContent
	}
	response.Json(listPerfiles, status, w)
}

func CrearPerfil(w http.ResponseWriter, r *http.Request) {
	var body models.Perfil
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
	e, perfil := services.CrearPerfil(body, permisos.Rol)
	if e != nil {
		response.Error(e.Error(), http.StatusBadRequest, w)
		return
	}
	response.Json(perfil, http.StatusOK, w)
}

func ActualizarPerfil(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var body models.Perfil
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
	e, perfil := services.ActualizarPerfil(body, vars, permisos.ID, permisos.Rol)
	if e != nil {
		response.Error(e.Error(), http.StatusBadRequest, w)
		return
	}
	response.Json(perfil, http.StatusOK, w)
}
func EliminarModuloPerfil(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var body models.Perfil
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
	log.Println("estoy Aqui", body)
	e, perfil := services.QuitarModulosPerfil(body, vars, permisos.ID, permisos.Rol)
	if e != nil {
		response.Error(e.Error(), http.StatusBadRequest, w)
		return
	}
	response.Json(perfil, http.StatusOK, w)
}
func EliminarPerfil(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	err, permisos := ConsultarUsuarioPerfil(r)
	if err != nil {
		response.Error(err.Error(), http.StatusBadRequest, w)
		return
	}
	e := services.EliminarPerfil(vars, permisos.ID, permisos.Rol)
	if e != nil {
		response.Error(e.Error(), http.StatusBadRequest, w)
		return
	}
	response.Json("Perfil Eliminado Correctamente", http.StatusOK, w)
}
