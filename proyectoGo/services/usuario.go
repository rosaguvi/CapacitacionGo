package services

import (
	"errors"
	"regexp"
	"strconv"

	"adminModPerl/crypto"
	"adminModPerl/database"
	"adminModPerl/models"
)

func ListarUsuarios(vars map[string]string, usu, id_rol uint) (lisUsuario []models.Usuario, err error) {
	if id_rol != 1 && id_rol != 2 {
		err = errors.New("El Rol del usuario no esta permitido para este modulo..")
		return
	}
	if id, ok := vars["id"]; ok {
		usuarioId, e := strconv.Atoi(id)
		if e != nil {
			err = e
			return
		}
		var usuario models.Usuario
		if id_rol == 1 {
			usuario = database.GetUsuario(uint(usuarioId))
		} else {
			usuario = database.GetUsuarioUsu(uint(usuarioId), usu)
		}
		lisUsuario = append(lisUsuario, usuario)
	} else {
		if id_rol == 1 {
			lisUsuario = database.GetUsuarios()

		} else {
			lisUsuario = database.GetUsuariosUsu(usu)
		}
	}
	return
}

func CrearUsuario(usuario models.Usuario, id_rol uint) (err error, nuevoUsuario models.Usuario) {
	if id_rol != 1 && id_rol != 2 {
		err = errors.New("El Rol del usuario no esta permitido para este modulo..")
		return
	}
	if id_rol == 1 {
		usuario.RolId = 2
	} else {
		usuario.RolId = 3
	}
	err = validarUsuario(usuario)
	if err != nil {
		return
	}
	nuevoUsuario = database.GetUserByEmail(usuario.Email)
	if nuevoUsuario.ID > 0 {
		err = errors.New("El usuario ya existe..")
		return
	}
	pass, e := crypto.EncryptText(usuario.Password)
	if e != nil {
		err = e
		return
	}
	usuario.Password = pass
	err, nuevoUsuario = database.CrearUsuario(usuario)
	if nuevoUsuario.ID <= 0 {
		err = errors.New("Error al crear el Usuario: " + err.Error())
		return
	}
	return
}
func ActualizarUsuario(usuario models.Usuario, vars map[string]string, usu_gb, id_rol uint) (err error, nuevo_usuario models.Usuario) {
	if id_rol != 1 && id_rol != 2 {
		err = errors.New("El Rol del usuario no esta permitido para este modulo..")
		return
	}
	if id, ok := vars["id"]; ok {
		usuarioId, e := strconv.Atoi(id)
		if e != nil {
			err = e
			return
		}
		nuevo_usuario = database.GetUsuario(uint(usuarioId))
		if nuevo_usuario.ID == 0 {
			err = errors.New("Usuario no existe.. ")
			return
		}
		if id_rol != 1 {
			if nuevo_usuario.UsuarioGB != usu_gb {
				err = errors.New("No tiene permisos de Edición sobre este usuario.. ")
				return
			}
		}
		usuario.ID = uint(usuarioId)
		err, usuario = validarUsuarioAct(usuario)
		if err != nil {
			return
		}
		nuevo_usuario = database.EditarUsuario(uint(usuarioId), usuario)
		nuevo_usuario = database.GetUsuario(uint(usuarioId))
	} else {
		err = errors.New("Falta el ID del Usuario")
	}
	return
}
func EliminarUsuario(vars map[string]string, usu_gb, id_rol uint) (err error) {
	if id_rol != 1 && id_rol != 2 {
		err = errors.New("El Rol del usuario no esta permitido para este modulo..")
		return
	}
	if id, ok := vars["id"]; ok {
		usuarioId, e := strconv.Atoi(id)
		if e != nil {
			err = e
			return
		}
		if uint(usuarioId) == usu_gb {
			err = errors.New("No puedes Eliminar tu propio Usuario.. ")
			return
		}
		usuario := database.GetUsuario(uint(usuarioId))
		if usuario.ID == 0 {
			err = errors.New("Usuario no existe.. ")
			return
		}
		if id_rol != 1 {
			if usuario.UsuarioGB != usu_gb {
				err = errors.New("No tiene permisos de Edición sobre este usuario.. ")
				return
			}
		}
		database.EliminarUsuario(uint(usuarioId))
	} else {
		err = errors.New("Falta el ID del Usuario")
	}
	return
}

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
		pass, e := crypto.EncryptText(usuario.Password)
		if e != nil {
			err := e
			return err, usuario
		}
		usuario.Password = pass
	}
	if usuario.RolId >= 0 {
		usuario.RolId = 0
	}
	usu = usuario
	usu.ID = 0
	return
}
