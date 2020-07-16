package jwt

import (
	"time"

	"github.com/ArturoArreola/demo-social-network/models"
	jwt "github.com/dgrijalva/jwt-go"
)

// GenerarJWT es la función para poder generar el JSON Web Token
func GenerarJWT(t models.Usuario) (string, error) {

	miClave := []byte("not13n3!")
	payload := jwt.MapClaims{
		"email":            t.Email,
		"nombre":           t.Nombre,
		"apellidos":        t.Apellidos,
		"fecha_nacimiento": t.FechaNacimiento,
		"biografia":        t.Biografia,
		"ubicacion":        t.Ubicacion,
		"sitioweb":         t.SitioWeb,
		"_id":              t.ID.Hex(),
		"exp":              time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(miClave)
	if err != nil {
		return tokenStr, err
	}
	return tokenStr, nil

}
