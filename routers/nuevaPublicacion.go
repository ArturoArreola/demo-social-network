package routers

import (
	"encoding/json"
	"github.com/ArturoArreola/demo-social-network/bd"
	"github.com/ArturoArreola/demo-social-network/models"
	"net/http"
	"time"
)

// NuevaPublicacion es la función para registrar un nuevo post
func NuevaPublicacion (w http.ResponseWriter, r *http.Request) {

	var mensaje models.Post
	err := json.NewDecoder(r.Body).Decode(&mensaje)

	registro := models.GrabarPost{
		UserID:  IDUsuario,
		Mensaje: mensaje.Mensaje,
		Fecha:   time.Now(),
	}

	_, status, err := bd.NuevoPost(registro)
	if err != nil {
		http.Error(w, "Ocurrió un error al intentar insertar el registro, reintente nuevamente"+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "No se ha logrado insertar el Tweet", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
