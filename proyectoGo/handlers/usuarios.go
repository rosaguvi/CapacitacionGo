package handlers

import (
	"errors"
	"net/http"

	"adminModPerl/io/request"
	"adminModPerl/io/response"
	"adminModPerl/jwt"
	"adminModPerl/models"
	"adminModPerl/services"

	"github.com/gorilla/mux"
)

func ListarUsuarios(w http.ResponseWriter, r *http.Request) {
	status := http.StatusOK
	vars := mux.Vars(r)
	err, permisos := ConsultarUsuarioPerfil(r)
	if err != nil {
		response.Error(err.Error(), http.StatusBadRequest, w)
		return
	}
	listUsuarios, e := services.ListarUsuarios(vars, permisos.ID, permisos.Rol)
	if e != nil {
		response.Error(e.Error(), http.StatusBadRequest, w)
		return
	}
	if len(listUsuarios) <= 0 || listUsuarios[0].ID == 0 {
		status = http.StatusNoContent
	}
	response.Json(listUsuarios, status, w)
}
func ConsultarModulosUsuario(w http.ResponseWriter, r *http.Request) {
	status := http.StatusOK
	err, permisos := ConsultarUsuarioPerfil(r)
	if err != nil {
		response.Error(err.Error(), http.StatusBadRequest, w)
		return
	}
	listModulos := services.ListarModulosPerfil(permisos.Perfil)
	if len(listModulos) <= 0 {
		response.Error("El perfil del Usuario no tiene modulos configurados", http.StatusBadRequest, w)
		return
	}
	response.Json(listModulos, status, w)
}

func ConsultarUsuarioPerfil(r *http.Request) (err error, claims *models.Claim) {
	header := r.Header.Get("Authorization")
	claims, err = jwt.ProcessToken(header)
	if err != nil {
		err = errors.New("Error al Validar Permisos para el Usuario")
		return
	}
	return
}
func CrearUsuario(w http.ResponseWriter, r *http.Request) {
	var body models.Usuario
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
	e, usuario := services.CrearUsuario(body, permisos.Rol)
	if e != nil {
		response.Error(e.Error(), http.StatusBadRequest, w)
		return
	}
	response.Json(usuario, http.StatusOK, w)
}

func ActualizarUsuario(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var body models.Usuario
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
	e, usuario := services.ActualizarUsuario(body, vars, permisos.ID, permisos.Rol)
	if e != nil {
		response.Error(e.Error(), http.StatusBadRequest, w)
		return
	}
	response.Json(usuario, http.StatusOK, w)
}
func EliminarUsuario(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	err, permisos := ConsultarUsuarioPerfil(r)
	if err != nil {
		response.Error(err.Error(), http.StatusBadRequest, w)
		return
	}
	e := services.EliminarUsuario(vars, permisos.ID, permisos.Rol)
	if e != nil {
		response.Error(e.Error(), http.StatusBadRequest, w)
		return
	}
	response.Json("Usuario Eliminado Exitosamente", http.StatusOK, w)
}
