package routers

import (
	"io"
	"net/http"
	"os"

	"github.com/ArturoArreola/demo-social-network/bd"
)

// ObtenerBanner envía el avatar al HTTP
func ObtenerBanner(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Debe de enviar el parámetro ID", http.StatusBadRequest)
		return
	}

	perfil, err := bd.BuscarPerfil(ID)
	if err != nil {
		http.Error(w, "Usuario no encontrado", http.StatusBadRequest)
		return
	}

	Imagen, err := os.Open("uploads/banners/"+perfil.Banner)
	if err != nil {
		http.Error(w, "Avatar no encontrado", http.StatusBadRequest)
		return
	}

	_, err = io.Copy(w, Imagen)
	if err != nil {
		http.Error(w, "Error al copiar la imagen del banner", http.StatusBadRequest)
	}
}