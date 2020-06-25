package bd

import (
	"context"
	"time"

	"github.com/ArturoArreola/demo-social-network/models"
)

// BorrarRelacion elimina la relación en la base de datos
func BorrarRelacion (relacion models.Relacion) (bool, error) {

	ctx, cancel := context.WithTimeout(context.Background(),time.Second*15)
	defer cancel()

	db := MongoCN.Database("social-network-db")
	coleccion := db.Collection("relacion")

	_, err := coleccion.DeleteOne(ctx, relacion)
	if err != nil {
		return false, err
	}

	return true, nil
}
