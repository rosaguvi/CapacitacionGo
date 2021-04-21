package middlewares

import (
	"errors"
	"net/http"
	"strings"

	"adminModPerl/database"
	"adminModPerl/jwt"
)

func Autenticar(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get("Authorization")
		_, err := jwt.ProcessToken(header)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(err.Error()))
			return
		}
		next.ServeHTTP(w, r)
	}
}

func ValidarPerfil(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		modulo := strings.Split(r.URL.Path[1:], "/")
		header := r.Header.Get("Authorization")
		claims, _ := jwt.ProcessToken(header)
		//log.Println("perfil", claims.Perfil)
		//log.Println("Modulo", modulo)
		perfilUsuario := database.GetPerfilModulo(claims.Perfil, modulo[0])
		//log.Println("perfil", perfilUsuario)
		if len(perfilUsuario.Modulos) <= 0 {
			w.WriteHeader(http.StatusUnauthorized)
			err := errors.New("Usuario no Autorizado para este modulo...")
			w.Write([]byte(err.Error()))
			return
		}
		next.ServeHTTP(w, r)
	}
}
