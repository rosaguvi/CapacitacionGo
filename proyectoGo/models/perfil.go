package models

import (
	"gorm.io/gorm"
)

type Perfil struct {
	gorm.Model
	Name      string   `json:"name"`
	Modulos   []Modulo `gorm:" many2many: periles_modulos; " json:"modulos,omitempty"`
	UsuarioGB uint     `json:"usuario_gb"`
}
