package bd

import(
	"context"
	"time"

	"github.com/ArturoArreola/demo-social-network/models"
)

// InsertarRelacion graba la relación en la base de datos
func InsertarRelacion(relacion models.Relacion) (bool, error){

	ctx, cancel := context.WithTimeout(context.Background(),time.Second*15)
	defer cancel()

	db := MongoCN.Database("social-network-db")
	coleccion := db.Collection("relacion")

	_, err := coleccion.InsertOne(ctx, relacion)
	if err != nil {
		return false, err
	}

	return true, nil

}
