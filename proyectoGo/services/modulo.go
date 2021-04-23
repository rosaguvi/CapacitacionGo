package services

import (
	"errors"
	"strconv"

	"adminModPerl/database"
	"adminModPerl/models"
)

/**
* Metodo que recibe los parametros y valida el perfil del usuario, para listar uno o varios Modulos
* @param map[string]string -- variables desde la dirección
* @param usu uint -- identificador del susuario logueado
* @param id_rol uint -- identificador del Rol del usuario
* @return []models.Modulo -- lista de obejeto Modulo
* @return error
**/
func ListarModulos(vars map[string]string, usu, id_rol uint) (lisModulo []models.Modulo, err error) {
	// valida si el rol del usuario esta autorizado para este modulo
	if id_rol != 1 && id_rol != 2 {
		err = errors.New("El Rol del usuario no esta permitido para este modulo..")
		return
	}
	// si se va a listar un solo modulo
	if id, ok := vars["id"]; ok {
		moduloId, e := strconv.Atoi(id)
		if e != nil {
			err = e
			return
		}
		var Modulo models.Modulo
		if id_rol == 1 {
			// si esta consultando el super admin
			Modulo = database.GetModulo(uint(moduloId))
		} else {
			// si esta consultando el admin
			Modulo = database.GetModuloUsu(uint(moduloId), usu)
		}
		lisModulo = append(lisModulo, Modulo)
		// si se va aconsultar todos los modulos
	} else {
		if id_rol == 1 {
			// si esta consultando el super admin
			lisModulo = database.GetModulos()

		} else {
			// si esta consultando el admin
			lisModulo = database.GetModulosUsu(usu)
		}
	}
	return
}

/**
* Consulta un perfil y reporna los modulos asociados a este
* @param id_rol uint -- identificador del Rol
* @return []models.Modulo -- lista de obejeto Modulo
**/
func ListarModulosPerfil(id_perfil uint) (lisModulos []models.Modulo) {
	perfil := database.GetPerfil(uint(id_perfil))
	lisModulos = perfil.Modulos
	return
}

/**
* Metodo que recibe los parametros y valida el perfil del usuario, para crear un modulo
* @param Objeto de tipo Modulo
* @param id_rol uint -- identificador del Rol del usuario
* @return Modulo -- obejeto Modulo creado
* @return error
**/
func CrearModulo(modulo models.Modulo, id_rol uint) (err error, nuevo_mod models.Modulo) {
	if id_rol != 1 && id_rol != 2 {
		err = errors.New("El Rol del usuario no esta permitido para este modulo..")
		return
	}
	// valida que todos los campos del objeto tengan la informacion correcta
	err = validarModulo(modulo)
	if err != nil {
		return
	}
	// valida que no exista otro modulo con ese nombre
	nuevo_mod = database.GetModuloNombre(modulo.Name)
	if nuevo_mod.ID > 0 {
		err = errors.New("El modulo ya Existe")
		return
	}
	// envia el objeto a la base de datos para ser creado
	err, nuevo_mod = database.CrearModulo(modulo)
	if nuevo_mod.ID <= 0 {
		err = errors.New("Error al crear el Modulo: " + err.Error())
		return
	}
	return
}

/**
* Metodo que recibe los parametros y valida el perfil del usuario, para actualizar un modulo
* @param Objeto de tipo Modulo
* @param map[string]string -- variables desde la dirección
* @param usu uint -- identificador del susuario logueado
* @param id_rol uint -- identificador del Rol del usuario
* @return Modulo -- obejeto Modulo actualizado
* @return error
**/
func ActualizarModulo(modulo models.Modulo, vars map[string]string, usu_gb, id_rol uint) (err error, nuevoModulo models.Modulo) {
	// valida si el rol enviado esta permitido para este modulo
	if id_rol != 1 && id_rol != 2 {
		err = errors.New("El Rol del usuario no esta permitido para este modulo..")
		return
	}
	// extrae el id del mosulo aactualizar para continuar con el proceso
	if id, ok := vars["id"]; ok {
		moduloId, e := strconv.Atoi(id)
		if e != nil {
			err = e
			return
		}
		// consulta el modulo por el id
		nuevoModulo = database.GetModulo(uint(moduloId))
		if nuevoModulo.ID == 0 {
			err = errors.New("Modulo no existe.. ")
			return
		}
		//si el rol es administrador y el usuario no es el que grabo el modulo no permite la actualizacion
		if id_rol != 1 {
			if nuevoModulo.UsuarioGB != usu_gb {
				err = errors.New("No tiene permisos de Edición sobre este Modulo.. ")
				return
			}
		}
		//valida que no exista otro modulo con el nombre enviado para actualizar
		nuevoModulo = database.GetModuloNombre(modulo.Name)
		if nuevoModulo.ID > 0 && nuevoModulo.ID != uint(moduloId) {
			err = errors.New("El nombre de Modulo ya esta asignado a otro Modulo")
			return
		}
		//Actualiza el modulo en la base de datos
		nuevoModulo = database.EditarModulo(uint(moduloId), modulo)
		//consulta el objeto actualizado
		nuevoModulo = database.GetModulo(uint(moduloId))
	} else {
		err = errors.New("Falta el ID del Modulo")
	}
	return
}

/**
* Metodo que recibe los parametros y valida el perfil del usuario, para eliminar un modulo
* @param map[string]string -- variables desde la dirección
* @param usu_gb uint -- identificador del susuario logueado
* @param id_rol uint -- identificador del Rol del usuario
* @return error
**/
func EliminarModulo(vars map[string]string, usu_gb, id_rol uint) (err error) {
	if id_rol != 1 && id_rol != 2 {
		err = errors.New("El Rol del usuario no esta permitido para este modulo..")
		return
	}
	// consulta el id del modulo a eliminar
	if id, ok := vars["id"]; ok {
		moduloId, e := strconv.Atoi(id)
		if e != nil {
			err = e
			return
		}
		// consulta el modulo por id, para validar que exista
		modulo := database.GetModulo(uint(moduloId))
		if modulo.ID == 0 {
			err = errors.New("Modulo no existe...")
			return
		}
		// si el rol es Admin, valida que el usuario sea quien creo el modulo
		if id_rol != 1 {
			if modulo.UsuarioGB != usu_gb {
				err = errors.New("No tiene permisos de Edición sobre este Modulo.. ")
				return
			}
		}
		// elimina el objeto de la base de datos
		database.EliminarModulo(uint(moduloId))
	} else {
		err = errors.New("Falta el ID del Modulo")
	}
	return
}

/**
* Metodo que valida que los campos tengan la informacion requerida
* @param Objeto de tipo Modulo
* @return error
**/
func validarModulo(mod models.Modulo) (err error) {
	if len(mod.Name) <= 0 {
		err = errors.New("Se debe ingresar un Nombre para el modulo")
		return
	}
	if mod.UsuarioGB == 0 {
		err = errors.New("Se debe asignar un usuario de creación")
		return
	}
	return
}
