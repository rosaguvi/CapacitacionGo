package database

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"adminModPerl/models"
)

var db *gorm.DB
var err error

const (
	host     = "localhost"
	port     = "5432"
	user     = "postgres"
	password = "mLemouw105"
	dbname   = "ProAdmModPer"
)

func GetConn() *gorm.DB {
	if db != nil {
		return db
	}
	conexion := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err = gorm.Open(postgres.Open(conexion), &gorm.Config{})
	if err != nil {
		log.Println("Error en la conexión...")
		panic(err)
	}
	return db
}

func Migrate() {
	_ = GetConn()
	db.AutoMigrate(&models.Usuario{}, &models.Modulo{}, &models.Perfil{}, &models.Rol{})
}

/**
* Metodo que inicializa los registros requeridos para el funcionamiento del aplicactivo
* @param null
* @return null
 */
func Inicializar() {
	rol := models.Rol{Name: "SUPER_ADMIN"}
	_, rol = CrearRol(rol)
	rol = models.Rol{Name: "ADMIN"}
	_, rol = CrearRol(rol)
	rol = models.Rol{Name: "USER"}
	_, rol = CrearRol(rol)

	perfil := models.Perfil{Name: "Jefe de Sistemas"}
	_, perfil = CrearPerfil(perfil)
	usuario := models.Usuario{Name: "Juan", Email: "Juan@personalsoft.com", Password: "$2a$08$nsfezcuZV0lsHJpIV7N2ZOHbbr/.zThlyP17Klm5Ypcjmu/b0PsUe", RolId: 1, PerfilId: perfil.ID}
	_, usuario = CrearUsuario(usuario)
	perfil.UsuarioGB = usuario.ID
	modulo := models.Modulo{Name: "Usuarios", UsuarioGB: usuario.ID}
	_, modulo = CrearModulo(modulo)
	var modulos []models.Modulo
	modulos = append(modulos, modulo)
	modulo = models.Modulo{Name: "Perfiles", UsuarioGB: usuario.ID}
	_, modulo = CrearModulo(modulo)
	modulos = append(modulos, modulo)
	modulo = models.Modulo{Name: "Modulos", UsuarioGB: usuario.ID}
	_, modulo = CrearModulo(modulo)
	modulos = append(modulos, modulo)
	perfil.Modulos = modulos
	perfil = EditarPerfil(perfil.ID, perfil)
}
