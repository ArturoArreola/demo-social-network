package bd

import (
	"context"
	"time"
	"fmt"

	"github.com/ArturoArreola/demo-social-network/models"
	"go.mongodb.org/mongo-driver/bson"
)

// ConsultaRelacion consulta la relación entre 2 usuarios
func ConsultaRelacion(relacion models.Relacion) (bool, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("social-network-db")
	coleccion := db.Collection("relacion")

	condicion := bson.M{
		"usuarioId": relacion.UsuarioID,
		"usuarioRelacionId": relacion.UsuarioRelacionID,
	}

	var resultado models.Relacion
	fmt.Println(resultado)

	err := coleccion.FindOne(ctx, condicion).Decode(&resultado)
	if err != nil {
		fmt.Println(err.Error())
		return false, err
	}

	return true, nil

}