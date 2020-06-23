package routers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"github.com/ArturoArreola/demo-social-network/bd"
)

// ObtenerPublicacion permite extraer los valores del Usuario
func ObtenerPublicacion (w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "No contiene el parámetro ID", http.StatusBadRequest)
		return
	}

	if len(r.URL.Query().Get("pagina")) < 1 {
		http.Error(w, "No contiene el parámetro página", http.StatusBadRequest)
		return
	}

	pagina, err := strconv.Atoi(r.URL.Query().Get("pagina"))

	if err != nil {
		http.Error(w, "Se debe enviar el parámetro página con valor mayor a 0", http.StatusBadRequest)
		return
	}

	pag := int64(pagina)
	respuesta, correcto := bd.BuscarPublicaciones(ID, pag)
	if correcto == false {
		http.Error(w, "Error al leer las publicaciones -> " + err.Error(), 400)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(respuesta)

}
