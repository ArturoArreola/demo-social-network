package middlew

import (
	"net/http"

	"github.com/ArturoArreola/demo-social-network/bd"
)

// CheckDB es el middleware que permite conocer el estado de conexión de la base de datos
func CheckDB(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if bd.CheckConnection() == 0 {
			http.Error(w, "Conexión perdida con la Base de Datos", 500)
			return
		}
		next.ServeHTTP(w, r)
	}
}
