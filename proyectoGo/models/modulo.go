package models

import (
	"gorm.io/gorm"
)

type Modulo struct {
	gorm.Model
	Name      string `json:"name"`
	UsuarioGB uint   `json:"usuario_gb"`
}
