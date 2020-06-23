package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/ArturoArreola/demo-social-network/bd"
	"github.com/ArturoArreola/demo-social-network/models"
)

// NuevaPublicacion
func NuevaPublicacion (w http.ResponseWriter, r *http.Request) {

	var mensaje models.Post
	err := json.NewDecoder(r.Body).Decode(&mensaje)

	registro := models.GrabarPost{
		UserID: IDUsuario,
		Mensaje: mensaje.Mensaje,
		Fecha: time.Now(),
	}

	_, status, err := bd.NuevoPost(registro)

	if err != nil {
		http.Error(w, "Ocurrió un error al intentar grabar el post -> " + err.Error(), http.StatusBadRequest)
		return
	}

	if status == false {
		http.Error(w, "no se pudo grabar el post", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)


}
