package services

import (
	"errors"
	"strconv"

	"adminModPerl/database"
	"adminModPerl/models"
)

func ListarModulos(vars map[string]string, usu, id_rol uint) (lisModulo []models.Modulo, err error) {
	if id_rol != 1 && id_rol != 2 {
		err = errors.New("El Rol del usuario no esta permitido para este modulo..")
		return
	}
	if id, ok := vars["id"]; ok {
		moduloId, e := strconv.Atoi(id)
		if e != nil {
			err = e
			return
		}
		var Modulo models.Modulo
		if id_rol == 1 {
			Modulo = database.GetModulo(uint(moduloId))
		} else {
			Modulo = database.GetModuloUsu(uint(moduloId), usu)
		}
		lisModulo = append(lisModulo, Modulo)
	} else {
		if id_rol == 1 {
			lisModulo = database.GetModulos()

		} else {
			lisModulo = database.GetModulosUsu(usu)
		}
	}
	return
}
func ListarModulosPerfil(id_perfil uint) (lisModulos []models.Modulo) {
	perfil := database.GetPerfil(uint(id_perfil))
	lisModulos = perfil.Modulos
	return
}

func CrearModulo(modulo models.Modulo, id_rol uint) (err error, nuevo_mod models.Modulo) {
	if id_rol != 1 && id_rol != 2 {
		err = errors.New("El Rol del usuario no esta permitido para este modulo..")
		return
	}
	err = validarModulo(modulo)
	if err != nil {
		return
	}
	nuevo_mod = database.GetModuloNombre(modulo.Name)
	if nuevo_mod.ID > 0 {
		err = errors.New("El modulo ya Existe")
		return
	}
	err, nuevo_mod = database.CrearModulo(modulo)
	if nuevo_mod.ID <= 0 {
		err = errors.New("Error al crear el Modulo: " + err.Error())
		return
	}
	return
}
func ActualizarModulo(modulo models.Modulo, vars map[string]string, usu_gb, id_rol uint) (err error, nuevoModulo models.Modulo) {
	if id_rol != 1 && id_rol != 2 {
		err = errors.New("El Rol del usuario no esta permitido para este modulo..")
		return
	}
	if id, ok := vars["id"]; ok {
		moduloId, e := strconv.Atoi(id)
		if e != nil {
			err = e
			return
		}
		nuevoModulo = database.GetModulo(uint(moduloId))
		if nuevoModulo.ID == 0 {
			err = errors.New("Modulo no existe.. ")
			return
		}
		if id_rol != 1 {
			if nuevoModulo.UsuarioGB != usu_gb {
				err = errors.New("No tiene permisos de Edición sobre este Modulo.. ")
				return
			}
		}
		nuevoModulo = database.GetModuloNombre(modulo.Name)
		if nuevoModulo.ID > 0 && nuevoModulo.ID != uint(moduloId) {
			err = errors.New("El nombre de Modulo ya esta asignado a otro Modulo")
			return
		}
		nuevoModulo = database.EditarModulo(uint(moduloId), modulo)
		nuevoModulo = database.GetModulo(uint(moduloId))
	} else {
		err = errors.New("Falta el ID del Modulo")
	}
	return
}
func EliminarModulo(vars map[string]string, usu_gb, id_rol uint) (err error) {
	if id_rol != 1 && id_rol != 2 {
		err = errors.New("El Rol del usuario no esta permitido para este modulo..")
		return
	}
	if id, ok := vars["id"]; ok {
		moduloId, e := strconv.Atoi(id)
		if e != nil {
			err = e
			return
		}
		modulo := database.GetModulo(uint(moduloId))
		if modulo.ID == 0 {
			err = errors.New("Modulo no existe...")
			return
		}
		if id_rol != 1 {
			if modulo.UsuarioGB != usu_gb {
				err = errors.New("No tiene permisos de Edición sobre este Modulo.. ")
				return
			}
		}
		database.EliminarModulo(uint(moduloId))
	} else {
		err = errors.New("Falta el ID del Modulo")
	}
	return
}

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
