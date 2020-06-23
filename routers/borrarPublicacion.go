package routers

import (
	"net/http"
	"github.com/ArturoArreola/demo-social-network/bd"
)

// BorrarPublicacion permite extraer los valores para poder eliminar el post
func BorrarPublicacion (w http.ResponseWriter, r *http.Request)  {

	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parámetro ID", http.StatusBadRequest)
		return
	}

	err := bd.BorrarPost(ID, IDUsuario)
	if err != nil {
		http.Error(w, "Ocurrió un error al intentar borrar el post" + err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)

}
