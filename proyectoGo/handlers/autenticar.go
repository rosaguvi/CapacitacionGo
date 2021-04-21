package handlers

import (
	"net/http"

	"golang.org/x/crypto/bcrypt"

	"adminModPerl/database"
	"adminModPerl/io/request"
	"adminModPerl/io/response"
	jwt "adminModPerl/jwt"
	"adminModPerl/models"
)

func Login(w http.ResponseWriter, r *http.Request) {
	var body models.Usuario
	err := request.Json(r, &body)
	if err != nil {
		response.Error("El formato del Body no es correcto", http.StatusBadRequest, w)
		return
	}
	if len(body.Email) <= 0 || len(body.Password) <= 0 {
		response.Error("El email y el password son obligatorios", http.StatusBadRequest, w)
		return
	}
	usuario := database.GetUserByEmail(body.Email)
	if usuario.ID <= 0 {
		response.Error("usuario no existe", http.StatusBadRequest, w)
		return
	}
	usuarioPass := []byte(usuario.Password)
	bodyPass := []byte(body.Password)
	err = bcrypt.CompareHashAndPassword(usuarioPass, bodyPass)
	if err != nil {
		response.Error("Error en el password", http.StatusBadRequest, w)
		return
	}
	token, err := jwt.GenerateToken(usuario)
	if err != nil {
		response.Error("Error generando el Token JWT", http.StatusInternalServerError, w)
		return
	}
	usuario.Password = ""
	loginResp := models.LoginResponse{Token: token, Usuario: usuario}
	response.Json(loginResp, http.StatusOK, w)
}
