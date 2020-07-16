package routers

import (
	"errors"
	"strings"

	"github.com/ArturoArreola/demo-social-network/bd"
	"github.com/ArturoArreola/demo-social-network/models"
	jwt "github.com/dgrijalva/jwt-go"
)

// Email es el email usado en todos los EndPoints
var Email string

// IDUsuario es el ID que se devuelve del modelo y que se va a usar en todos los EndPoints
var IDUsuario string

// ProcesarToken es la función donde se procesa el token para obtener sus valores
func ProcesarToken(tk string) (*models.Claim, bool, string, error) {

	miClave := []byte("not13n3!")
	claims := &models.Claim{}

	splitToken := strings.Split(tk, "Bearer")
	if len(splitToken) != 2 {
		return claims, false, string(""), errors.New("formato de token invalido")
	}

	tk = strings.TrimSpace(splitToken[1])

	tkn, err := jwt.ParseWithClaims(tk, claims, func(token *jwt.Token) (interface{}, error) {
		return miClave, nil
	})
	if err == nil {
		_, encontrado, _ := bd.CheckUsuarioExiste(claims.Email)
		if encontrado == true {
			Email = claims.Email
			IDUsuario = claims.ID.Hex()
		}
		return claims, encontrado, IDUsuario, nil
	}
	if !tkn.Valid {
		return claims, false, string(""), errors.New("token Inválido")
	}
	return claims, false, string(""), err
}
