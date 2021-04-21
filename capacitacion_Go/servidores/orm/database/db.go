package database

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"modCapacitacion/servidores/orm/models"
)

const (
	host     = "localhost"
	port     = "5432"
	user     = "postgres"
	password = "mLemouw105"
	dbname   = "Capacitacion_go"
)

var db *gorm.DB

var err error

func GetConn() *gorm.DB {
	if db != nil {
		return db
	}
	conexion := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err = gorm.Open(postgres.Open(conexion), &gorm.Config{})
	if err != nil {
		log.Println("Error en la conexion")
		panic(err)
	}
	return db
}

func Migrar() {
	_ = GetConn()
	db.AutoMigrate(&models.Curso{})
	log.Print("Migrando datos..")
}

func GetCursos() (listaCursos []models.Curso) {
	_ = GetConn()
	db.Find(&listaCursos)
	return
}

func CrearCurso(nombre, descripcion string) (curso models.Curso) {
	_ = GetConn()
	curso = models.Curso{Nombre: nombre, Descripcion: descripcion}
	db.Create(&curso)
	db.Last(&curso)
	return
}

func VerCurso(id uint) (curso models.Curso) {
	_ = GetConn()
	db.Find(&curso, id)
	return
}
func EditarCurso(id uint, nombre, descripcion string) (curso models.Curso) {
	_ = GetConn()
	db.Find(&curso, id)
	// db.Model(&curso).Update("Nombre", nombre)                                                    // Actualizar campos individuales
	// db.Model(&curso).Update("Descrpcion", descripcion)                                           // Actualizar campos individuales
	// db.Model(&curso).Updates(models.Curso{Nombre: nombre, Descripcion: descripcion})             // Actualizar todas las propiedades
	db.Model(&curso).Updates(map[string]interface{}{"Nombre": nombre, "Descripcion": descripcion}) // Actualizar solo las propiedades que están en el mapa
	return
}

func EliminarCurso(id uint) {
	_ = GetConn()
	db.Delete(&models.Curso{}, id)
}
