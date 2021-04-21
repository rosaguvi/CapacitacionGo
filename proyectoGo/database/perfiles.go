package database

import (
	"adminModPerl/models"
)

func GetPerfiles() (listaPerfiles []models.Perfil) {
	_ = GetConn()
	db.Preload("Modulos").Find(&listaPerfiles)
	return
}
func GetPerfilesUsu(id_gra uint) (listaPerfiles []models.Perfil) {
	_ = GetConn()
	db.Preload("Modulos").Where("usuario_gb = ? ", id_gra).Find(&listaPerfiles)
	return
}

func CrearPerfil(perfilini models.Perfil) (error, models.Perfil) {
	_ = GetConn()
	err := db.Create(&perfilini)
	if err != nil {
		return err.Error, perfilini
	}
	db.Last(&perfilini)
	return nil, perfilini
}
func GetPerfil(id uint) (perfil models.Perfil) {
	_ = GetConn()
	db.Preload("Modulos").Find(&perfil, id)
	return
}
func GetPerfilUsu(id, id_gra uint) (perfil models.Perfil) {
	_ = GetConn()
	db.Preload("Modulos").Where("usuario_gb = ? ", id_gra).Find(&perfil, id)
	return
}
func GetPerfilModulo(id uint, mombreModulo string) (perfil models.Perfil) {
	_ = GetConn()
	db.Preload("Modulos", "name ilike ?", mombreModulo).Find(&perfil, id)
	return
}
func GetPerfilNombre(nombre string) (perfil models.Perfil) {
	_ = GetConn()
	db.Where("name ilike ? ", nombre).Find(&perfil)
	return
}

func EditarPerfil(id uint, perfilIni models.Perfil) (perfil models.Perfil) {
	_ = GetConn()
	db.Find(&perfil, id)
	db.Model(&perfil).Updates(perfilIni) // Actualizar todas las propiedades
	db.Model(&perfil).Association("Modulos").Append(perfilIni.Modulos)
	return
}
func QuitarModulosPerfil(id uint, perfilIni models.Perfil) (perfil models.Perfil) {
	_ = GetConn()
	db.Find(&perfil, id)
	db.Model(&perfil).Association("Modulos").Delete(perfilIni.Modulos)
	return
}

func EliminarPerfil(id uint) {
	_ = GetConn()
	db.Delete(&models.Perfil{}, id)
}
