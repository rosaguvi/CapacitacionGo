package database

import (
	"adminModPerl/models"
)

func GetUsuarios() (listaUsuarios []models.Usuario) {
	_ = GetConn()
	db.Preload("Perfil").Preload("Rol").Find(&listaUsuarios)
	return
}

// Retorna el listado de Usuarios creados por un usuario
func GetUsuariosUsu(id_gra uint) (listaUsuarios []models.Usuario) {
	_ = GetConn()
	db.Preload("Perfil").Preload("Rol").Where("usuario_gb = ? ", id_gra).Find(&listaUsuarios)
	return
}

func CrearUsuario(usuarioini models.Usuario) (error, models.Usuario) {
	_ = GetConn()
	err := db.Create(&usuarioini)
	if err != nil {
		return err.Error, usuarioini
	}
	db.Last(&usuarioini)
	return nil, usuarioini
}

func GetUsuario(id uint) (usuario models.Usuario) {
	_ = GetConn()
	db.Preload("Perfil").Preload("Rol").Find(&usuario, id)
	return
}

func GetUsuarioUsu(id, id_gra uint) (usuario models.Usuario) {
	_ = GetConn()
	db.Preload("Perfil").Preload("Rol").Where("usuario_gb = ? ", id_gra).Find(&usuario, id)
	return
}

func EditarUsuario(id uint, usuarioIni models.Usuario) (usuario models.Usuario) {
	_ = GetConn()
	db.Find(&usuario, id)
	db.Model(&usuario).Updates(usuarioIni) // Actualizar todas las propiedades
	return
}

func EliminarUsuario(id uint) {
	_ = GetConn()
	db.Delete(&models.Usuario{}, id)
}

func GetUserByEmail(email string) (usuario models.Usuario) {
	_ = GetConn()
	db.Find(&usuario, "email = ?", email)
	return
}
