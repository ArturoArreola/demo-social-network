package bd

import (
	"fmt"

	"github.com/ArturoArreola/demo-social-network/models"
	"golang.org/x/crypto/bcrypt"
)

// IntentoLogin es la función con la cual verifico que el usuario exista y que la contraseña sea la correcta
func IntentoLogin(email string, password string) (models.Usuario, bool) {
	usuario, encontrado, idUsuario := CheckUsuarioExiste(email)

	fmt.Println("Este es el Id del Usuario %s", idUsuario)
	if encontrado == false {
		return usuario, false
	}

	passwordBytes := []byte(password)
	passwordBD := []byte(usuario.Password)

	err := bcrypt.CompareHashAndPassword(passwordBD, passwordBytes)

	if err != nil {
		return usuario, false
	}
	return usuario, true

}
