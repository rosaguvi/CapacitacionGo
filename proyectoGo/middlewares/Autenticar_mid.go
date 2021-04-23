package middlewares

import (
	"errors"
	"net/http"
	"strings"

	"adminModPerl/database"
	"adminModPerl/jwt"
)

/**
* Metodo que valida si el token enviado en la cabecera es correcto
* para autorizar el ingreso a la aplicacion
* @param http.ResponseWriter
* @param *http.Request
* @return http.ResponseWriter
* @return *http.Request
**/
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

/**
* Metodo que valida si el usuario esta autorizado a usar el programa que esta consultando
* y autorizar el acceso al mismo
* @param http.ResponseWriter
* @param *http.Request
* @return http.ResponseWriter
* @return *http.Request
**/
func ValidarPerfil(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		modulo := strings.Split(r.URL.Path[1:], "/")
		header := r.Header.Get("Authorization")
		// obtiene los datos de logueo del usuario desde el token
		claims, _ := jwt.ProcessToken(header)
		//busca el perfil del usuario y ademas consulta el modulo solicitado en la direccion.
		perfilUsuario := database.GetPerfilModulo(claims.Perfil, modulo[0])
		// valaida si el modulo que se desea acceder esta en el perfil del usuario.
		if len(perfilUsuario.Modulos) <= 0 {
			w.WriteHeader(http.StatusUnauthorized)
			err := errors.New("Usuario no Autorizado para este modulo...")
			w.Write([]byte(err.Error()))
			return
		}
		next.ServeHTTP(w, r)
	}
}
