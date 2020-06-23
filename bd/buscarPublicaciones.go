package bd

import (
	"context"
	"log"
	"time"

	"github.com/ArturoArreola/demo-social-network/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// BuscarPublicaciones
func BuscarPublicaciones (ID string, pagina int64) ([]*models.ObtenerPosts, bool) {

	ctx, cancel := context.WithTimeout(context.Background(),time.Second*15)
	defer cancel()

	db := MongoCN.Database("social-network-db")
	coleccion := db.Collection("tweet")

	var resultado []*models.ObtenerPosts

	condicion := bson.M{
		"userid":ID,
	}

	opciones := options.Find()
	opciones.SetLimit(20)
	opciones.SetSort(bson.D{{Key: "fecha", Value: -1}})		// Va a traer todos los valores ordenados por fecha en orden descendente (esto último se indica con el -1)
	opciones.SetSkip((pagina -1) * 20)

	cursor, err := coleccion.Find(ctx, condicion, opciones)

	if err != nil {
		log.Fatal(err.Error())
		return resultado, false
	}

	for cursor.Next(context.TODO()){
		var registro models.ObtenerPosts
		err := cursor.Decode(&registro)
		if err != nil {
			return resultado, false
		}

		resultado = append(resultado, &registro)
	}

	return resultado, true

}