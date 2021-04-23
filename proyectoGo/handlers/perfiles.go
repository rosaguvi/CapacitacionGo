package handlers

import (
	"net/http"

	"github.com/gorilla/mux"

	"adminModPerl/io/request"
	"adminModPerl/io/response"
	"adminModPerl/models"
	"adminModPerl/services"
)

/**
* Metodo que extrae la variables y consulta los permisos del usuario y
* envia esta información al servce para listar uno o varios Perfiles
* @param http.ResponseWriter
* @param *http.Request
* @return http.ResponseWriter
**/
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

/**
* Metodo que extrae la variables y consulta los permisos del usuario y
* envia esta información al servce para crear un nuevo perfil en la base de datos
* @param http.ResponseWriter
* @param *http.Request
* @return http.ResponseWriter
**/
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

/**
* Metodo que extrae la variables y consulta los permisos del usuario y
* envia esta información al servce para Actualizar un perfil
* @param http.ResponseWriter
* @param *http.Request
* @return http.ResponseWriter
**/
func ActualizarPerfil(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var body models.Perfil
	// consulta datos de rol, perfil e ID desde el token del usuario
	err, permisos := ConsultarUsuarioPerfil(r)
	if err != nil {
		response.Error(err.Error(), http.StatusBadRequest, w)
		return
	}
	//Extrae el objeto que viene en el cuerpo del mensaje
	err = request.Json(r, &body)
	if err != nil {
		response.Error("El Formato del Body no es el Correcto", http.StatusBadRequest, w)
		return
	}
	// Envia los parametros a service para la valdiacion y acutializacion del registro
	e, perfil := services.ActualizarPerfil(body, vars, permisos.ID, permisos.Rol)
	if e != nil {
		response.Error(e.Error(), http.StatusBadRequest, w)
		return
	}
	response.Json(perfil, http.StatusOK, w)
}

/**
* Metodo que extrae la variables y consulta los permisos del usuario y
* envia esta información al servce para desasociar uno a varios Modulos de un Perfil
* @param http.ResponseWriter
* @param *http.Request
* @return http.ResponseWriter
**/
func EliminarModuloPerfil(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var body models.Perfil
	// consulta datos de rol, perfil e ID desde el token del usuario
	err, permisos := ConsultarUsuarioPerfil(r)
	if err != nil {
		response.Error(err.Error(), http.StatusBadRequest, w)
		return
	}
	//Extrae el objeto que viene en el cuerpo del mensaje
	err = request.Json(r, &body)
	if err != nil {
		response.Error("El Formato del Body no es el Correcto", http.StatusBadRequest, w)
		return
	}
	// Envia la información al service para realizar las valdiaciones y desasociar los modulos del perfil
	e, perfil := services.QuitarModulosPerfil(body, vars, permisos.ID, permisos.Rol)
	if e != nil {
		response.Error(e.Error(), http.StatusBadRequest, w)
		return
	}
	response.Json(perfil, http.StatusOK, w)
}

/**
* Metodo que extrae la variables y consulta los permisos del usuario y
* envia esta información al servce para Eliminar un perfil
* @param http.ResponseWriter
* @param *http.Request
* @return http.ResponseWriter
**/
func EliminarPerfil(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	// consulta datos de rol, perfil e ID desde el token del usuario
	err, permisos := ConsultarUsuarioPerfil(r)
	if err != nil {
		response.Error(err.Error(), http.StatusBadRequest, w)
		return
	}
	// Envia la información al service para realizar las valdiaciones y Eliminar el perfil
	e := services.EliminarPerfil(vars, permisos.ID, permisos.Rol)
	if e != nil {
		response.Error(e.Error(), http.StatusBadRequest, w)
		return
	}
	response.Json("Perfil Eliminado Correctamente", http.StatusOK, w)
}
