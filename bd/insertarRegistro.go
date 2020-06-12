package bd

import (
	"context"
	"time"

	"github.com/ArturoArreola/demo-social-network/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func InsertarRegistro(usuario models.Usuario) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("social-network-db")
	coleccion := db.Collection("usuarios")

	usuario.Password, _ = EncriptarPassword(usuario.Password)

	result, err := coleccion.InsertOne(ctx, usuario)

	if err != nil {
		return "", false, err
	}

	// Esta es la manera de obtener el ID de un registro que acaba de insertar en MongoDB
	ObjID, _ := result.InsertedID.(primitive.ObjectID)
	return ObjID.String(), true, nil
}
