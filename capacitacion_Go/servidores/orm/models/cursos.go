package models

import (
	// paquetes Externos
	"gorm.io/gorm"
)

type Curso struct {
	gorm.Model
	Nombre      string
	Descripcion string
}
