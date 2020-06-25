package routers

import (
	"encoding/json"
	"net/http"

	"github.com/ArturoArreola/demo-social-network/bd"
	"github.com/ArturoArreola/demo-social-network/models"
)

// ConsultaRelacion checa si hay relación entre 2 usuarios
func ConsultaRelacion(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")

	var relacion models.Relacion
	relacion.UsuarioID = IDUsuario
	relacion.UsuarioRelacionID = ID

	var respuesta models.RespuestaRelacion

	status, err := bd.ConsultaRelacion(relacion)
	if err != nil || status == false {
		respuesta.Status = false
	} else {
		respuesta.Status = true
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(respuesta)

}
