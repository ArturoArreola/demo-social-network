package routers

import (
	"io"
	"net/http"
	"strings"
	"os"

	"github.com/ArturoArreola/demo-social-network/bd"
	"github.com/ArturoArreola/demo-social-network/models"
)

// SubirBanner sube la imagen del banner al servidor
func SubirBanner(w http.ResponseWriter, r *http.Request) {

	file, handler, err := r.FormFile("banner")
	var extension = strings.Split(handler.Filename, ".")[1]
	var archivo string = "uploads/banners/banner-" + IDUsuario + "." + extension

	f, err := os.OpenFile(archivo, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		http.Error(w, "Error al subir la imagen -> " + err.Error(),400)
		return
	}

	_, err = io.Copy(f, file)
	if err != nil {
		http.Error(w, "Error al copiar la imagen -> " + err.Error(),400)
		return
	}

	var usuario models.Usuario
	var status bool

	usuario.Banner = "banner-" + IDUsuario + "." + extension
	status, err2 := bd.ModificarRegistro(usuario, IDUsuario)
	if err2 != nil || status == false {
		http.Error(w, "Error al guardar el banner en la base de datos-> " + err2.Error(),400)
		return
	}

	w.Header().Set("content-type","application/json")
	w.WriteHeader(http.StatusOK)

}
