package routers

import (
	"net/http"

	"github.com/ArturoArreola/demo-social-network/models"
	"github.com/ArturoArreola/demo-social-network/bd"
)

// AltaRelacion realiza el registro de la relacion entre usuarios
func AltaRelacion(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "El parámetro ID es obligatorio", http.StatusBadRequest)
		return
	}

	var relacion models.Relacion
	relacion.UsuarioID = IDUsuario
	relacion.UsuarioRelacionID = ID

	estatus, err := bd.InsertarRelacion(relacion)

	if err != nil {
		http.Error(w, "Error al intentar insertar en la base de datos -> " + err.Error(), http.StatusBadRequest)
		return
	}

	if estatus == false {
		http.Error(w, "No se logró insertar la relación -> " + err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)

}
