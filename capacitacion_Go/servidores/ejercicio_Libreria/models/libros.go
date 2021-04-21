package models

import (
	// paquetes Externos
	"gorm.io/gorm"
)

type Libro struct {
	gorm.Model
	Titulo          string
	Descripcion     string
	Genero          string
	Autor           string
	Editorial       string
	A_publicacion   int
	NumCopVendiadas int
}
