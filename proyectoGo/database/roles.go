package database

import (
	"adminModPerl/models"
)

func GetRoles() (listaRoles []models.Rol) {
	_ = GetConn()
	db.Find(&listaRoles)
	return
}

func CrearRol(rolini models.Rol) (error, models.Rol) {
	_ = GetConn()
	err := db.Create(&rolini)
	if err != nil {
		return err.Error, rolini
	}
	db.Last(&rolini)
	return nil, rolini
}

func GetRol(id uint) (rol models.Rol) {
	_ = GetConn()
	db.Find(&rol, id)
	return
}

func EditarRol(id uint, rolIni models.Rol) (rol models.Rol) {
	_ = GetConn()
	db.Find(&rol, id)
	db.Model(&rol).Updates(rolIni) // Actualizar todas las propiedades
	return
}

func EliminarRol(id uint) {
	_ = GetConn()
	db.Delete(&models.Rol{}, id)
}
