package services

import (
	"errors"
	"regexp"
	"strconv"

	"adminModPerl/crypto"
	"adminModPerl/database"
	"adminModPerl/models"
)

/**
* Metodo que recibe los parametros y valida el perfil del usuario,
* para listar uno o varios Usuarios
* @param map[string]string -- variables desde la dirección
* @param usu uint -- identificador del susuario logueado
* @param id_rol uint -- identificador del Rol del usuario
* @return []Usuario -- listado de obejetos de tipo usuario
* @return error
**/
func ListarUsuarios(vars map[string]string, usu, id_rol uint) (lisUsuario []models.Usuario, err error) {
	// valida si el Rol esta habilitado para este modulo
	if id_rol != 1 && id_rol != 2 {
		err = errors.New("El Rol del usuario no esta permitido para este modulo..")
		return
	}
	// extrade el id del registro a consulta
	if id, ok := vars["id"]; ok {
		usuarioId, e := strconv.Atoi(id)
		if e != nil {
			err = e
			return
		}
		var usuario models.Usuario
		/*
			dependiendo del rol consulta el usuario por ID de registro o por ID de registro
			y de usuario que grabo
		*/
		if id_rol == 1 {
			usuario = database.GetUsuario(uint(usuarioId))
		} else {
			usuario = database.GetUsuarioUsu(uint(usuarioId), usu)
		}
		lisUsuario = append(lisUsuario, usuario)
	} else {
		/*
			dependiendo del rol consulta el listado de todos los usuarios o
			El listado de los usuarios grabados por un usuario especifico
		*/
		if id_rol == 1 {
			lisUsuario = database.GetUsuarios()

		} else {
			lisUsuario = database.GetUsuariosUsu(usu)
		}
	}
	return
}

/**
* Metodo que recibe los parametros y valida el perfil del usuario,
* para crear un Usuario
* @param Objeto de tipo Usuario
* @param usu uint -- identificador del susuario logueado
* @param id_rol uint -- identificador del Rol del usuario
* @return Usuario -- obejeto de tipo usuario
* @return error
**/
func CrearUsuario(usuario models.Usuario, id_rol uint) (err error, nuevoUsuario models.Usuario) {
	// valida si el Rol esta habilitado para este modulo
	if id_rol != 1 && id_rol != 2 {
		err = errors.New("El Rol del usuario no esta permitido para este modulo..")
		return
	}
	// asigna el rol del registro de acuerdo al rol del usuario logueado
	if id_rol == 1 {
		usuario.RolId = 2
	} else {
		usuario.RolId = 3
	}
	// valida que el registro tenga la informacion correcta
	err = validarUsuario(usuario)
	if err != nil {
		return
	}
	// valdia que no exista otr usuario con el email enviado
	nuevoUsuario = database.GetUserByEmail(usuario.Email)
	if nuevoUsuario.ID > 0 {
		err = errors.New("El usuario ya existe..")
		return
	}
	// encripta la clave para el usuario
	pass, e := crypto.EncryptText(usuario.Password)
	if e != nil {
		err = e
		return
	}
	usuario.Password = pass
	// crea el usuario y si hay error le da manejo
	err, nuevoUsuario = database.CrearUsuario(usuario)
	if nuevoUsuario.ID <= 0 {
		err = errors.New("Error al crear el Usuario: " + err.Error())
		return
	}
	return
}

/**
* Metodo que recibe los parametros y valida el perfil del usuario,
* para actualizar un Usuario
* @param Objeto de tipo Usuario
* @param map[string]string -- variables desde la dirección
* @param usu_gb uint -- identificador del susuario logueado
* @param id_rol uint -- identificador del Rol del usuario
* @return Usuario -- obejeto de tipo usuario
* @return error
**/
func ActualizarUsuario(usuario models.Usuario, vars map[string]string, usu_gb, id_rol uint) (err error, nuevo_usuario models.Usuario) {
	// valida si el Rol esta habilitado para este modulo
	if id_rol != 1 && id_rol != 2 {
		err = errors.New("El Rol del usuario no esta permitido para este modulo..")
		return
	}
	// extrae el id del usuario a procesar
	if id, ok := vars["id"]; ok {
		usuarioId, e := strconv.Atoi(id)
		if e != nil {
			err = e
			return
		}
		// valdia que el usuario exista
		nuevo_usuario = database.GetUsuario(uint(usuarioId))
		if nuevo_usuario.ID == 0 {
			err = errors.New("Usuario no existe.. ")
			return
		}
		// valdia los permisos del usuario sobre el registro
		if id_rol != 1 {
			if nuevo_usuario.UsuarioGB != usu_gb {
				err = errors.New("No tiene permisos de Edición sobre este usuario.. ")
				return
			}
		}
		// valdia que los campos que se actualizaran tenga la informacion correcta
		usuario.ID = uint(usuarioId)
		err, usuario = validarUsuarioAct(usuario)
		if err != nil {
			return
		}
		// actualiza el usuario
		nuevo_usuario = database.EditarUsuario(uint(usuarioId), usuario)
		// consulta el usuario actualizado
		nuevo_usuario = database.GetUsuario(uint(usuarioId))
	} else {
		err = errors.New("Falta el ID del Usuario")
	}
	return
}

/**
* Metodo que recibe los parametros y valida el perfil del usuario,
* para eliminar un Usuario
* @param map[string]string -- variables desde la dirección
* @param usu_gb uint -- identificador del susuario logueado
* @param id_rol uint -- identificador del Rol del usuario
* @return Usuario -- obejeto de tipo usuario
* @return error
**/
func EliminarUsuario(vars map[string]string, usu_gb, id_rol uint) (err error) {
	// valida si el Rol esta habilitado para este modulo
	if id_rol != 1 && id_rol != 2 {
		err = errors.New("El Rol del usuario no esta permitido para este modulo..")
		return
	}
	// EXtrae el id del registro a eliminar
	if id, ok := vars["id"]; ok {
		usuarioId, e := strconv.Atoi(id)
		if e != nil {
			err = e
			return
		}
		// valida para no permitir que el usuario elimine su propio usuario
		if uint(usuarioId) == usu_gb {
			err = errors.New("No puedes Eliminar tu propio Usuario.. ")
			return
		}
		// valdia que el registro exista
		usuario := database.GetUsuario(uint(usuarioId))
		if usuario.ID == 0 {
			err = errors.New("Usuario no existe.. ")
			return
		}
		// valdia los permisos del usuario sobre el registro
		if id_rol != 1 {
			if usuario.UsuarioGB != usu_gb {
				err = errors.New("No tiene permisos de Edición sobre este usuario.. ")
				return
			}
		}
		// Elimina el usuario
		database.EliminarUsuario(uint(usuarioId))
	} else {
		err = errors.New("Falta el ID del Usuario")
	}
	return
}

/**
* Metodo que valida que los campos tengan la informacion requerida
* @param Objeto de tipo Perfil
* @return error
**/
func validarUsuario(usuario models.Usuario) (err error) {
	if len(usuario.Name) < 5 {
		err = errors.New("El nombre debe contener minimo 5 Caracteres")
		return
	}
	regex := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	if len(usuario.Email) < 6 || len(usuario.Email) > 255 || !regex.MatchString(usuario.Email) {
		err = errors.New("el email no tiene el formato adecuado")
		return
	}
	if len(usuario.Password) < 6 {
		err = errors.New("El Password del Usuario debe contener minimo 6 caracteres")
		return
	}
	if usuario.PerfilId <= 0 {
		err = errors.New("Se debe Asignar un perfil al usuario")
		return
	}
	if usuario.RolId <= 0 {
		err = errors.New("Se debe Asignar un Rol al usuario")
		return
	}
	return
}

/**
* Metodo que valida que los campos que se actualizaran tengan la informacion requerida
* @param Objeto de tipo Perfil
* @return error
**/
func validarUsuarioAct(usuario models.Usuario) (err error, usu models.Usuario) {
	if usuario.Name != "" && len(usuario.Name) < 5 {
		err = errors.New("El nombre debe contener minimo 5 Caracteres")
		return err, usuario
	}
	regex := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	if usuario.Email != "" {
		if len(usuario.Email) < 6 || len(usuario.Email) > 255 || !regex.MatchString(usuario.Email) {
			err = errors.New("el email no tiene el formato adecuado")
			return err, usuario
		}
		// si se va a actualizar le correo valida que el correo no pertenezca a otro usuario
		usu = database.GetUserByEmail(usuario.Email)
		if usu.ID > 0 && usu.ID != usuario.ID {
			err = errors.New("El email ya esta asociado a otro usuario ..")
			return
		}
	}
	if usuario.Password != "" {
		if len(usuario.Password) < 6 {
			err := errors.New("El Password del Usuario debe contener minimo 6 caracteres")
			return err, usuario
		}
		// si el password se va actualizar se encripta
		pass, e := crypto.EncryptText(usuario.Password)
		if e != nil {
			err := e
			return err, usuario
		}
		usuario.Password = pass
	}
	// si se va actualizar el rol lo deja en cero para que n se actualice.
	if usuario.RolId >= 0 {
		usuario.RolId = 0
	}
	usu = usuario
	usu.ID = 0
	return
}
