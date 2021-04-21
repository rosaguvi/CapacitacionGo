package jwt

import (
	"errors"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"

	"adminModPerl/models"
)

var JWT_KEY = []byte("RoAgudelo2105@PersonalSoft.Clave")

// UserEmail keep user email
var UserEmail string

// UserID keep user ID
var UserID string

// GenerateToken returns jwt token or error
func GenerateToken(user models.Usuario) (string, error) {
	payload := jwt.MapClaims{
		"email":  user.Email,
		"name":   user.Name,
		"id":     user.ID,
		"perfil": user.PerfilId,
		"rol":    user.RolId,
		"exp":    time.Now().Add(time.Hour * 24).Unix(),
		// email, roles...
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, payload).SignedString(JWT_KEY)
	return token, err
}

// ProcessToken check if token is valid
func ProcessToken(authHeader string) (*models.Claim, error) {
	claims := &models.Claim{}
	if len(authHeader) <= 0 {
		return claims, errors.New("La penticion no Contiene el Token")
	}
	splitToken := strings.Split(authHeader, "Bearer")
	if len(splitToken) != 2 {
		return claims, errors.New("El formato del Token no es valido")
	}
	tokenStr := strings.TrimSpace(splitToken[1])
	_, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return JWT_KEY, nil
	})
	if err != nil {
		return claims, err
	}
	return claims, nil
}
