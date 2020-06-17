package middlew

import (
	"net/http"

	"github.com/ArturoArreola/demo-social-network/routers"
)

// ValidarJWT es la función que permite validar el JWT que viene en la petición
func ValidarJWT(next http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		_, _, _, err := routers.ProcesarToken(r.Header.Get("Authorization"))

		if err != nil {
			http.Error(w, "Error al validar el Token "+err.Error(), http.StatusBadRequest)
			return
		}
		next.ServeHTTP(w, r)
	}
}
