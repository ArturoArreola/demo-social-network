package routers

import (
	"encoding/json"
	"net/http"
	"github.com/ArturoArreola/demo-social-network/bd"
)

// VerPerfil permite extraer los valores del Perfil del Usuario
func VerPerfil(w http.ResponseWriter, r *http.Request){

	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "No contiene el parámetro ID", http.StatusBadRequest)
		return
	}

	perfil, err := bd.BuscarPerfil(ID)
	if err != nil {
		http.Error(w, "No se encontró el ID en la base de datos -> " + err.Error(), 400)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(perfil)
}
