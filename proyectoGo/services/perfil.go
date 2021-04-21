package services

import (
	"errors"
	"log"
	"strconv"

	"adminModPerl/database"
	"adminModPerl/models"
)

func ListarPerfiles(vars map[string]string, usu, id_rol uint) (lisPerfil []models.Perfil, err error) {
	if id_rol != 1 && id_rol != 2 {
		err = errors.New("El Rol del usuario no esta permitido para este modulo..")
		return
	}
	if id, ok := vars["id"]; ok {
		PerfilId, e := strconv.Atoi(id)
		if e != nil {
			err = e
			return
		}
		var Perfil models.Perfil
		if id_rol == 1 {
			Perfil = database.GetPerfil(uint(PerfilId))
		} else {
			Perfil = database.GetPerfilUsu(uint(PerfilId), usu)
		}
		lisPerfil = append(lisPerfil, Perfil)
	} else {
		if id_rol == 1 {
			lisPerfil = database.GetPerfiles()

		} else {
			lisPerfil = database.GetPerfilesUsu(usu)
		}
	}
	return
}
func CrearPerfil(peril models.Perfil, id_rol uint) (err error, nuevo_Pefil models.Perfil) {
	if id_rol != 1 && id_rol != 2 {
		err = errors.New("El Rol del usuario no esta permitido para este modulo..")
		return
	}
	err = validarPerfil(peril)
	if err != nil {
		return
	}
	nuevo_Pefil = database.GetPerfilNombre(peril.Name)
	//log.Println("modulo Existe", nuevo_Pefil)
	if nuevo_Pefil.ID > 0 {
		err = errors.New("El Perfil ya Existe")
		return
	}
	err, nuevo_Pefil = database.CrearPerfil(peril)
	if nuevo_Pefil.ID <= 0 {
		err = errors.New("Error al crear el Perfil: " + err.Error())
		return
	}
	return
}
func ActualizarPerfil(peril models.Perfil, vars map[string]string, usu_gb, id_rol uint) (err error, nuevoPerfil models.Perfil) {
	if id_rol != 1 && id_rol != 2 {
		err = errors.New("El Rol del usuario no esta permitido para este modulo..")
		return
	}
	if id, ok := vars["id"]; ok {
		perfilId, e := strconv.Atoi(id)
		if e != nil {
			err = e
			return
		}
		nuevoPerfil = database.GetPerfil(uint(perfilId))
		if nuevoPerfil.ID == 0 {
			err = errors.New("Perfil no existe.. ")
			return
		}
		if id_rol != 1 {
			if nuevoPerfil.UsuarioGB != usu_gb {
				err = errors.New("No tiene permisos de Edición sobre este Perfil.. ")
				return
			}
		}
		nuevoPerfil = database.EditarPerfil(uint(perfilId), peril)
		nuevoPerfil = database.GetPerfil(uint(perfilId))
	} else {
		err = errors.New("Falta el ID del Perfil")
	}
	return
}
func QuitarModulosPerfil(perfil models.Perfil, vars map[string]string, usu_gb, id_rol uint) (err error, nuevoPerfil models.Perfil) {
	if id_rol != 1 && id_rol != 2 {
		err = errors.New("El Rol del usuario no esta permitido para este modulo..")
		return
	}
	if id, ok := vars["id"]; ok {
		perfilId, e := strconv.Atoi(id)
		if e != nil {
			err = e
			return
		}
		nuevoPerfil = database.GetPerfil(uint(perfilId))
		if nuevoPerfil.ID == 0 {
			err = errors.New("Perfil no existe.. ")
			return
		}
		if id_rol != 1 {
			if nuevoPerfil.UsuarioGB != usu_gb {
				err = errors.New("No tiene permisos de Edición sobre este Perfil.. ")
				return
			}
		}
		log.Println("perfil antes de actualizar: ", perfil)
		nuevoPerfil = database.QuitarModulosPerfil(uint(perfilId), perfil)
		nuevoPerfil = database.GetPerfil(uint(perfilId))
		log.Println("perfil despues de actualizar: ", nuevoPerfil)
	} else {
		err = errors.New("Falta el ID del Perfil")
	}
	return
}
func EliminarPerfil(vars map[string]string, usu_gb, id_rol uint) (err error) {
	if id_rol != 1 && id_rol != 2 {
		err = errors.New("El Rol del usuario no esta permitido para este modulo..")
		return
	}
	if id, ok := vars["id"]; ok {
		perfilId, e := strconv.Atoi(id)
		if e != nil {
			err = e
			return
		}
		perfil := database.GetPerfil(uint(perfilId))
		if perfil.ID == 0 {
			err = errors.New("Perfil no existe.. ")
			return
		}
		if id_rol != 1 {
			if perfil.UsuarioGB != usu_gb {
				err = errors.New("No tiene permisos de Edición sobre este Perfil.. ")
				return
			}
		}
		database.EliminarPerfil(uint(perfilId))
	} else {
		err = errors.New("Falta el ID del Perfil")
	}
	return
}

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
