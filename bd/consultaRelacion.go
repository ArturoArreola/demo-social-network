package bd

import (
	"context"
	"github.com/ArturoArreola/demo-social-network/models"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

// ConsultaRelacion consulta la relación entre 2 usuarios
func ConsultaRelacion(relacion models.Relacion) (bool, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("social-network-db")
	coleccion := db.Collection("relacion")

	condicion := bson.M{
		"usuarioid":         relacion.UsuarioID,
		"usuariorelacionid": relacion.UsuarioRelacionID,
	}

	var resultado models.Relacion
	err := coleccion.FindOne(ctx, condicion).Decode(&resultado)
	if err != nil {
		return false, err
	}
	return true, nil
}
