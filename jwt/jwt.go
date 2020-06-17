package jwt

import (
	"time"

	"github.com/ArturoArreola/demo-social-network/models"
	jwt "github.com/dgrijalva/jwt-go"
)

// GenerarJWT es la función para poder generar el JSON Web Token
func GenerarJWT(usuario models.Usuario) (string, error) {

	miClave := []byte("not13n3!")
	payload := jwt.MapClaims{
		"email":            usuario.Email,
		"nombre":           usuario.Nombre,
		"apellidos":        usuario.Apellidos,
		"fecha_nacimiento": usuario.FechaNacimiento,
		"biografia":        usuario.Biografia,
		"ubicacion":        usuario.Ubicacion,
		"sitioweb":         usuario.SitioWeb,
		"_id":              usuario.ID.Hex(),
		"exp":              time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(miClave)

	if err != nil {
		return tokenStr, err
	}
	return tokenStr, nil

}
