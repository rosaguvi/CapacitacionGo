package models

import (
	"gorm.io/gorm"
)

type Rol struct {
	gorm.Model
	Name string `json:"name"`
}
