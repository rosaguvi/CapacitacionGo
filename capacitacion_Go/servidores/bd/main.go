package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = "5432"
	user     = "postgres"
	password = "mLemouw105"
	dbname   = "Capacitacion_go"
)

type Categoria struct {
	id          int
	nombre      string
	Descripcion string
}

var db *sql.DB
var err error

func main() {
	insertarCategoria("N categoria", "Nueva Categoria")
	listarCategoria()
}

func conectar() {
	conexion := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err = sql.Open("postgres", conexion)
	if err != nil {
		log.Println("Error en la conexion")
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		log.Println("Error en el ping")
		panic(err)
	}
	log.Println("conexión abierta")
}

func listarCategoria() {
	conectar()
	defer func() {
		db.Close()
		log.Println("conexion Cerrada")
	}()

	query := "select id, nombre, descripcion FROM categoria"
	rows, err := db.Query(query)
	if err != nil {
		log.Println("Error al consultar")
		panic(err)
	}
	listaCat := []Categoria{}

	for rows.Next() {
		cat := Categoria{}
		rows.Scan(
			&cat.id,
			&cat.nombre,
			&cat.Descripcion,
		)
		listaCat = append(listaCat, cat)
	}
	log.Println(listaCat)
}

func insertarCategoria(nombre, descripcion string) {
	conectar()
	defer func() {
		db.Close()
		log.Println("conexion Cerrada")
	}()

	query := fmt.Sprintf("INSERT INTO public.categoria(nombre, descripcion)	VALUES ('%s', '%s');", nombre, descripcion)
	_, err := db.Query(query)
	if err != nil {
		log.Println("Error al insertar en Categoria")
		panic(err)
	}

}
