package models

import (
	"gorm.io/gorm"
)

type Usuario struct {
	gorm.Model
	Name      string `json:"name"`
	Email     string `json:"email"`
	Password  string `json:"password,omitempty"`
	PerfilId  uint   `json:"perfil_id"`
	Perfil    Perfil `gorm:"foreignKey:PerfilId" json:"perfil_usuario,omitempty"`
	RolId     uint   `json:"rol_id"`
	Rol       Rol    `gorm:"foreignKey:RolId" json:"rol_usuario,omitempty"`
	UsuarioGB uint   `json:"usuario_gb"`
}
