package routers

import (
	"io"
	"net/http"
	"strings"
	"os"

	"github.com/ArturoArreola/demo-social-network/bd"
	"github.com/ArturoArreola/demo-social-network/models"
)

// SubirAvatar sube la imagen de avatar al servidor
func SubirAvatar(w http.ResponseWriter, r *http.Request) {

	file, handler, err := r.FormFile("avatar")
	var extension = strings.Split(handler.Filename, ".")[1]
	var archivo string = "uploads/avatars/avatar-" + IDUsuario + "." + extension

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

	usuario.Avatar = "avatar-" + IDUsuario + "." + extension
	status, err2 := bd.ModificarRegistro(usuario, IDUsuario)
	if err2 != nil || status == false {
		http.Error(w, "Error al guardar el avatar en la base de datos-> " + err2.Error(),400)
		return
	}

	w.Header().Set("content-type","application/json")
	w.WriteHeader(http.StatusOK)

}
