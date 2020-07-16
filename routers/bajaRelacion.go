package routers

import (
	"net/http"

	"github.com/ArturoArreola/demo-social-network/models"
	"github.com/ArturoArreola/demo-social-network/bd"
)

// BajaRelacion realiza la eliminación de la relacion entre usuarios
func BajaRelacion(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")
	var t models.Relacion
	t.UsuarioID = IDUsuario
	t.UsuarioRelacionID = ID

	estatus, err := bd.BorrarRelacion(t)

	if err != nil {
		http.Error(w, "Error al intentar borrar en la base de datos -> " + err.Error(), http.StatusBadRequest)
		return
	}

	if estatus == false {
		http.Error(w, "No se logró borrar la relación -> " + err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}
