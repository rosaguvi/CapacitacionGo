package database

import (
	"adminModPerl/models"
)

func GetModulos() (listaModulos []models.Modulo) {
	_ = GetConn()
	db.Preload("Perfil").Find(&listaModulos)
	return
}
func GetModulosUsu(id_gra uint) (listaModulos []models.Modulo) {
	_ = GetConn()
	db.Preload("Perfil").Where("usuario_gb = ? ", id_gra).Find(&listaModulos)
	return
}

func CrearModulo(moduloini models.Modulo) (error, models.Modulo) {
	_ = GetConn()
	err := db.Create(&moduloini)
	if err != nil {
		return err.Error, moduloini
	}
	db.Last(&moduloini)
	return nil, moduloini
}

func GetModulo(id uint) (modulo models.Modulo) {
	_ = GetConn()
	db.Preload("Perfil").Preload("Rol").Find(&modulo, id)
	return
}
func GetModuloNombre(nombre string) (modulo models.Modulo) {
	_ = GetConn()
	db.Where("name ilike ? ", nombre).Find(&modulo)
	return
}
func GetModuloUsu(id, id_gra uint) (modulo models.Modulo) {
	_ = GetConn()
	db.Preload("Perfil").Preload("Rol").Where("usuario_gb = ? ", id_gra).Find(&modulo, id)
	return
}

func EditarModulo(id uint, moduloIni models.Modulo) (modulo models.Modulo) {
	_ = GetConn()
	db.Find(&modulo, id)
	db.Model(&modulo).Updates(moduloIni) // Actualizar todas las propiedades
	return
}

func EliminarModulo(id uint) {
	_ = GetConn()
	db.Delete(&models.Modulo{}, id)
}
