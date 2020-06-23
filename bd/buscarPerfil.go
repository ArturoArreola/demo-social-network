package bd

import (
	"context"
	"fmt"
	"time"
	"github.com/ArturoArreola/demo-social-network/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// BuscarPerfil es la funcion para buscar el perfil del usuario con base en su ID usuario
func BuscarPerfil (ID string) (models.Usuario, error) {

	ctx, cancel := context.WithTimeout(context.Background(),time.Second*15)
	defer cancel()

	db := MongoCN.Database("social-network-db")
	coleccion := db.Collection("usuarios")

	var perfil  models.Usuario
	objID, _ := primitive.ObjectIDFromHex(ID)

	condicion := bson.M{
		"_id":objID,
	}

	err := coleccion.FindOne(ctx, condicion).Decode(&perfil)
	perfil.Password=""

	if err != nil {
		fmt.Println("Registro no encontrado" + err.Error())
		return perfil, err
	}

	return perfil, nil
}

