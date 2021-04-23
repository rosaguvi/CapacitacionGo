package services

import (
	"errors"
	"strconv"

	"adminModPerl/database"
	"adminModPerl/models"
)

/**
* Metodo que recibe los parametros y valida el perfil del usuario y los parametros enviados
* para listar uno o varios perfiles
* @param map[string]string -- variables desde la dirección
* @param usu uint -- identificador del susuario logueado
* @param id_rol uint -- identificador del Rol del usuario
* @return [] Perfil-- listados de obejetos de tipo PErfil
* @return error
**/
func ListarPerfiles(vars map[string]string, usu, id_rol uint) (lisPerfil []models.Perfil, err error) {
	// valida si el Rol esta habilitado para este modulo
	if id_rol != 1 && id_rol != 2 {
		err = errors.New("El Rol del usuario no esta permitido para este modulo..")
		return
	}
	// si hay id Perfil consultara solo uno
	if id, ok := vars["id"]; ok {
		PerfilId, e := strconv.Atoi(id)
		if e != nil {
			err = e
			return
		}
		var Perfil models.Perfil
		// si el rol es super
		if id_rol == 1 {
			Perfil = database.GetPerfil(uint(PerfilId))
		} else {
			Perfil = database.GetPerfilUsu(uint(PerfilId), usu)
		}
		lisPerfil = append(lisPerfil, Perfil)
	} else {
		// si no hay id perfil listara todos
		if id_rol == 1 {
			lisPerfil = database.GetPerfiles()
		} else {
			// si el rol es admin listara solo los que este haya creado
			lisPerfil = database.GetPerfilesUsu(usu)
		}
	}
	return
}

/**
* Metodo que recibe los parametros y valida el perfil del usuario y los parametros enviados
* para crear un perfil
* @param Objeto de tipo Perfil
* @param id_rol uint -- identificador del Rol del usuario
* @return Perfil-- obejeto de tipo Perfil
* @return error
**/
func CrearPerfil(peril models.Perfil, id_rol uint) (err error, nuevo_Pefil models.Perfil) {
	// valida si el Rol esta habilitado para este modulo
	if id_rol != 1 && id_rol != 2 {
		err = errors.New("El Rol del usuario no esta permitido para este modulo..")
		return
	}
	//valida que los campos necesarios tengan la informacion correcta
	err = validarPerfil(peril)
	if err != nil {
		return
	}
	//consulta si ya hay un perfil con ese nombre
	nuevo_Pefil = database.GetPerfilNombre(peril.Name)
	if nuevo_Pefil.ID > 0 {
		err = errors.New("El Perfil ya Existe")
		return
	}
	// envia acrear el perfil y si hay error le da manejo
	err, nuevo_Pefil = database.CrearPerfil(peril)
	if nuevo_Pefil.ID <= 0 {
		err = errors.New("Error al crear el Perfil: " + err.Error())
		return
	}
	return
}

/**
* Metodo que recibe los parametros y valida el perfil del usuario y los parametros enviados
* para actualizar un perfil
* @param Objeto de tipo Perfil
* @param map[string]string -- variables desde la dirección
* @param usu_gb uint -- identificador del susuario logueado
* @param id_rol uint -- identificador del Rol del usuario
* @return Perfil-- obejeto de tipo Perfil
* @return error
**/
func ActualizarPerfil(peril models.Perfil, vars map[string]string, usu_gb, id_rol uint) (err error, nuevoPerfil models.Perfil) {
	// valida si el Rol esta habilitado para este modulo
	if id_rol != 1 && id_rol != 2 {
		err = errors.New("El Rol del usuario no esta permitido para este modulo..")
		return
	}
	// extrae el id del perfil a actualizar
	if id, ok := vars["id"]; ok {
		perfilId, e := strconv.Atoi(id)
		if e != nil {
			err = e
			return
		}
		// valida que el perfil exista
		nuevoPerfil = database.GetPerfil(uint(perfilId))
		if nuevoPerfil.ID == 0 {
			err = errors.New("Perfil no existe.. ")
			return
		}
		// valida los permisos del usuario
		if id_rol != 1 {
			if nuevoPerfil.UsuarioGB != usu_gb {
				err = errors.New("No tiene permisos de Edición sobre este Perfil.. ")
				return
			}
		}
		// Actualiza el perfil
		nuevoPerfil = database.EditarPerfil(uint(perfilId), peril)
		// consulta el perfl actualizado
		nuevoPerfil = database.GetPerfil(uint(perfilId))
	} else {
		err = errors.New("Falta el ID del Perfil")
	}
	return
}

/**
* Metodo que recibe los parametros y valida el perfil del usuario y los parametros enviados
* para desasociar modulos de un perfril
* @param Objeto de tipo Perfil
* @param map[string]string -- variables desde la dirección
* @param usu_gb uint -- identificador del susuario logueado
* @param id_rol uint -- identificador del Rol del usuario
* @return Perfil-- obejetos de tipo Perfil
* @return error
**/
func QuitarModulosPerfil(perfil models.Perfil, vars map[string]string, usu_gb, id_rol uint) (err error, nuevoPerfil models.Perfil) {
	// valida si el Rol esta habilitado para este modulo
	if id_rol != 1 && id_rol != 2 {
		err = errors.New("El Rol del usuario no esta permitido para este modulo..")
		return
	}
	// extrade el id del peril a actualizar
	if id, ok := vars["id"]; ok {
		perfilId, e := strconv.Atoi(id)
		if e != nil {
			err = e
			return
		}
		// valdia que el perfil exista
		nuevoPerfil = database.GetPerfil(uint(perfilId))
		if nuevoPerfil.ID == 0 {
			err = errors.New("Perfil no existe.. ")
			return
		}
		// valdia los permisos del usuario
		if id_rol != 1 {
			if nuevoPerfil.UsuarioGB != usu_gb {
				err = errors.New("No tiene permisos de Edición sobre este Perfil.. ")
				return
			}
		}
		// envia el objeto a la base de datos para desasociar los modulos
		nuevoPerfil = database.QuitarModulosPerfil(uint(perfilId), perfil)
		//consulta el perfil actualizado
		nuevoPerfil = database.GetPerfil(uint(perfilId))
	} else {
		err = errors.New("Falta el ID del Perfil")
	}
	return
}

/**
* Metodo que recibe los parametros y valida el perfil del usuario y los parametros enviados
* para eliminar un perfil
* @param map[string]string -- variables desde la dirección
* @param usu_gb uint -- identificador del susuario logueado
* @param id_rol uint -- identificador del Rol del usuario
* @return error
**/
func EliminarPerfil(vars map[string]string, usu_gb, id_rol uint) (err error) {
	// valida si el Rol esta habilitado para este modulo
	if id_rol != 1 && id_rol != 2 {
		err = errors.New("El Rol del usuario no esta permitido para este modulo..")
		return
	}
	// extrade el id del perfil a actualizar
	// si existe hace el proceso
	if id, ok := vars["id"]; ok {
		perfilId, e := strconv.Atoi(id)
		if e != nil {
			err = e
			return
		}
		// valdia que el perfil exista
		perfil := database.GetPerfil(uint(perfilId))
		if perfil.ID == 0 {
			err = errors.New("Perfil no existe.. ")
			return
		}
		// valdia los permisos del usuario sobre el registro
		if id_rol != 1 {
			if perfil.UsuarioGB != usu_gb {
				err = errors.New("No tiene permisos de Edición sobre este Perfil.. ")
				return
			}
		}
		// elimina el pefil de la base de datos
		database.EliminarPerfil(uint(perfilId))
	} else {
		err = errors.New("Falta el ID del Perfil")
	}
	return
}

/**
* Metodo que valida que los campos tengan la informacion requerida
* @param Objeto de tipo Perfil
* @return error
**/
func validarPerfil(perfil models.Perfil) (err error) {
	if len(perfil.Name) <= 0 {
		err = errors.New("Se debe ingresar un Nombre para el Perfil")
		return
	}
	if perfil.UsuarioGB == 0 {
		err = errors.New("Se debe asignar un usuario de creación")
		return
	}
	return
}
