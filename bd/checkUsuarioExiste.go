package bd

import (
	"context"
	"time"

	"github.com/ArturoArreola/demo-social-network/models"
	"go.mongodb.org/mongo-driver/bson"
)

// CheckUsuarioExiste es la función que recibe un email de parámetro y revisa si ya existe el usuario
func CheckUsuarioExiste(email string) (models.Usuario, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	defer cancel()
	db := MongoCN.Database("social-network-db")
	coleccion := db.Collection("usuarios")

	condicion := bson.M{"email": email}

	var resultado models.Usuario

	err := coleccion.FindOne(ctx, condicion).Decode(&resultado)
	ID := resultado.ID.Hex()
	if err != nil {
		return resultado, false, ID
	}
	return resultado, true, ID
}
