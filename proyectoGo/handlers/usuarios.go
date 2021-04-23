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

/**
* Metodo que extrae las variables y consulta los permisos del usuario y
* envia esta información al servce Listar uno o varios usuarios
* @param http.ResponseWriter
* @param *http.Request
* @return http.ResponseWriter
**/
func ListarUsuarios(w http.ResponseWriter, r *http.Request) {
	status := http.StatusOK
	vars := mux.Vars(r)
	// consulta datos de rol, perfil e ID desde el token del usuario
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

/**
* Metodo que extrae las variables y consulta los permisos del usuario y
* envia esta información al servce Listar los modulos asociados al perfil del usuario
* @param http.ResponseWriter
* @param *http.Request
* @return http.ResponseWriter
**/
func ConsultarModulosUsuario(w http.ResponseWriter, r *http.Request) {
	status := http.StatusOK
	// consulta datos de rol, perfil e ID desde el token del usuario
	err, permisos := ConsultarUsuarioPerfil(r)
	if err != nil {
		response.Error(err.Error(), http.StatusBadRequest, w)
		return
	}
	//Envia el perfil del token, para consultar los modulos asociados al este
	listModulos := services.ListarModulosPerfil(permisos.Perfil)
	if len(listModulos) <= 0 {
		response.Error("El perfil del Usuario no tiene modulos configurados", http.StatusBadRequest, w)
		return
	}
	response.Json(listModulos, status, w)
}

/**
* Metodo que consulta los parametros de l usuario contenidos en el token y los retorna
* mediante el objeto claims
* @param *http.Request
* @return objeto de tipo *models.Claim
**/
func ConsultarUsuarioPerfil(r *http.Request) (err error, claims *models.Claim) {
	header := r.Header.Get("Authorization")
	claims, err = jwt.ProcessToken(header)
	if err != nil {
		err = errors.New("Error al Validar Permisos para el Usuario")
		return
	}
	return
}

/**
* Metodo que extrae las variables y consulta los permisos del usuario y
* envia esta información al service para crear un usuario
* @param http.ResponseWriter
* @param *http.Request
* @return http.ResponseWriter
**/
func CrearUsuario(w http.ResponseWriter, r *http.Request) {
	var body models.Usuario
	// consulta datos de rol, perfil e ID desde el token del usuario
	err, permisos := ConsultarUsuarioPerfil(r)
	if err != nil {
		response.Error(err.Error(), http.StatusBadRequest, w)
		return
	}
	// Extrae el objeto contenido en el cuerpo del mensaje
	err = request.Json(r, &body)
	if err != nil {
		response.Error("El Formato del Body no es el Correcto", http.StatusBadRequest, w)
		return
	}
	// asigna el usuario que graba desde la informacion contenida en el token
	body.UsuarioGB = permisos.ID
	// Envia los parametros necesarios al service para validar y crear el usuario
	e, usuario := services.CrearUsuario(body, permisos.Rol)
	if e != nil {
		response.Error(e.Error(), http.StatusBadRequest, w)
		return
	}
	response.Json(usuario, http.StatusOK, w)
}

/**
* Metodo que extrae las variables y consulta los permisos del usuario y
* envia esta información al service para actualizar un usuario
* @param http.ResponseWriter
* @param *http.Request
* @return http.ResponseWriter
**/
func ActualizarUsuario(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var body models.Usuario
	// consulta datos de rol, perfil e ID desde el token del usuario
	err, permisos := ConsultarUsuarioPerfil(r)
	if err != nil {
		response.Error(err.Error(), http.StatusBadRequest, w)
		return
	}
	// Extrae el objeto contenido en el cuerpo del mensaje
	err = request.Json(r, &body)
	if err != nil {
		response.Error("El Formato del Body no es el Correcto", http.StatusBadRequest, w)
		return
	}
	// Envia los parametros necesarios al service para validar y actualizar el usuario
	e, usuario := services.ActualizarUsuario(body, vars, permisos.ID, permisos.Rol)
	if e != nil {
		response.Error(e.Error(), http.StatusBadRequest, w)
		return
	}
	response.Json(usuario, http.StatusOK, w)
}

/**
* Metodo que extrae las variables y consulta los permisos del usuario y
* envia esta información al service para Eliminar un usuario
* @param http.ResponseWriter
* @param *http.Request
* @return http.ResponseWriter
**/
func EliminarUsuario(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	// consulta datos de rol, perfil e ID desde el token del usuario
	err, permisos := ConsultarUsuarioPerfil(r)
	if err != nil {
		response.Error(err.Error(), http.StatusBadRequest, w)
		return
	}
	// Envia los parametros necesarios al service para validar y Eliminar el usuario
	e := services.EliminarUsuario(vars, permisos.ID, permisos.Rol)
	if e != nil {
		response.Error(e.Error(), http.StatusBadRequest, w)
		return
	}
	response.Json("Usuario Eliminado Exitosamente", http.StatusOK, w)
}
