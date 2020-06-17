package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/ArturoArreola/demo-social-network/bd"
	"github.com/ArturoArreola/demo-social-network/jwt"
	"github.com/ArturoArreola/demo-social-network/models"
)

//	Login es la función para realizar el inicio de sesión
func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	var usuario models.Usuario

	err := json.NewDecoder(r.Body).Decode(&usuario)

	if err != nil {
		http.Error(w, "El usuario y/o contraseña son incorrectos"+err.Error(), 400)
		return
	}

	if len(usuario.Email) == 0 {
		http.Error(w, "El email es un campo requerido", 400)
		return
	}

	documento, existe := bd.IntentoLogin(usuario.Email, usuario.Password)

	if existe == false {
		http.Error(w, "El usuario y/o contraseña son incorrectos", 400)
		return
	}

	jwtKey, err := jwt.GenerarJWT(documento)

	if err != nil {
		http.Error(w, "Ocurrió un error"+err.Error(), 500)
		return
	}

	respuesta := models.RespuestaLogin{
		Token: jwtKey,
	}

	w.Header().Set("Content-Type", "aplication/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(respuesta)

	//	Manera de grabar una cookie desde el backend en Go
	expirationTime := time.Now().Add(24 * time.Hour)
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   jwtKey,
		Expires: expirationTime,
	})
}
