package database

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"modCapacitacion/servidores/ejercicio_Libreria/models"
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
	db.AutoMigrate(&models.Libro{})
	log.Print("Migrando datos..")
}

func GetLibros() (listaLibros []models.Libro) {
	_ = GetConn()
	db.Find(&listaLibros)
	return
}

func CrearLibro(libroini models.Libro) models.Libro {
	_ = GetConn()
	db.Create(&libroini)
	db.Last(&libroini)
	return libroini
}

func VerLibro(id uint) (libro models.Libro) {
	_ = GetConn()
	db.Find(&libro, id)
	return
}

func ConsultarLibroXAutor(autor string) (libro []models.Libro) {
	_ = GetConn()
	db.Where("autor ilike ?", "%"+autor+"%").Find(&libro)
	return
}
func ConsultarLibroXEditorial(editorial string) (libro []models.Libro) {
	_ = GetConn()
	log.Print(editorial)
	db.Where("editorial ilike ?", "%"+editorial+"%").Find(&libro)
	return
}
func ConLibrosAPrango(rango1, rango2 int) (libro []models.Libro) {
	_ = GetConn()
	db.Where("A_publicacion BETWEEN ? and ?", rango1, rango2).Find(&libro)
	return
}
func ConLibrosApublicacion(apublicacion int) (libro []models.Libro) {
	_ = GetConn()
	db.Where("A_publicacion = ? ", apublicacion).Find(&libro)
	return
}
func ConLibrosNumCopVendidas(num_copias int) (libro []models.Libro) {
	_ = GetConn()
	db.Where("NumCopVendiadas >= ? ", num_copias).Find(&libro)
	return
}
func EditarLibro(id uint, libroIni models.Libro) (libro models.Libro) {
	_ = GetConn()
	db.Find(&libro, id)
	db.Model(&libro).Updates(libroIni) // Actualizar todas las propiedades
	return
}

func EliminarLibro(id uint) {
	_ = GetConn()
	db.Delete(&models.Libro{}, id)
}
