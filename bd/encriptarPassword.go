package bd

import "golang.org/x/crypto/bcrypt"

// EncriptarPassword es la función para encriptar las contraseñas del usuario
func EncriptarPassword(password string) (string, error) {
	costo := 8
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), costo)
	return string(bytes), err
}
