package bd

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

// BorrarPost es la funcion para poder borrar un post de la base de datos
func BorrarPost (PostID string, UserID string) error {

	ctx, cancel := context.WithTimeout(context.Background(),time.Second*15)
	defer cancel()

	db := MongoCN.Database("social-network-db")
	coleccion := db.Collection("tweet")

	objID, _ := primitive.ObjectIDFromHex(PostID)

	condicion := bson.M{
		"_id": 		objID,
		"userid": 	UserID,
	}

	_, err := coleccion.DeleteOne(ctx, condicion)

	return err
}
