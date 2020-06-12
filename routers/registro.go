package routers

import (
	"encoding/json"
	"net/http"

	"github.com/ArturoArreola/demo-social-network/bd"
	"github.com/ArturoArreola/demo-social-network/models"
)

// Registro es la función para crear en la BD el registro del usuario
func Registro(w http.ResponseWriter, r *http.Request) {

	var usuario models.Usuario
	err := json.NewDecoder(r.Body).Decode(&usuario)

	if err != nil {
		http.Error(w, "Error en los datos recibidos", 400)
		return
	}

	if len(usuario.Email) == 0 {
		http.Error(w, "El email del usuario es requerido", 400)
		return
	}

	if len(usuario.Password) < 6 {
		http.Error(w, "La contraseña debe tener al menos 6 caracteres", 400)
		return
	}

	_, encontrado, _ := bd.CheckUsuarioExiste(usuario.Email)
	if encontrado == true {
		http.Error(w, "Ya existe un usuario registrado con este email", 400)
		return
	}

	_, estatus, err := bd.InsertarRegistro(usuario)
	if err != nil {
		http.Error(w, "Ocurrió un error al intentar registrar al usuario -> "+err.Error(), 400)
		return
	}

	if estatus == false {
		http.Error(w, "No se pudo dar de alta al usuario", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
