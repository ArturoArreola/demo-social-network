package routers

import (
	"encoding/json"
	"net/http"

	"github.com/ArturoArreola/demo-social-network/bd"
	"github.com/ArturoArreola/demo-social-network/models"
)

// ModificarPerfil es la función para poder modificar la información del perfil del usuario
func ModificarPerfil(w http.ResponseWriter, r *http.Request) {

	var modelo models.Usuario

	err := json.NewDecoder(r.Body).Decode(&modelo)
	if err != nil {
		http.Error(w, "Datos incorrectos -> " + err.Error(), 400)
		return
	}

	estatus, err2 := bd.ModificarRegistro(modelo, IDUsuario)
	if err2 !=nil {
		http.Error(w, "Ocurrió un error al intentar modificar el registro -> " + err.Error(), 400)
		return
	}

	if estatus == false {
		http.Error(w, "No se ha logrado modificar la información del usuario.", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
