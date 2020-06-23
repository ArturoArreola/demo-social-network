package bd

import(
	"context"
	"time"

	"github.com/ArturoArreola/demo-social-network/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// NuevoPost es la función para grabar el post en la base de datos
func NuevoPost(publicacion models.GrabarPost) (string, bool, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("social-network-db")
	coleccion := db.Collection("tweet")

	registro := bson.M{
		"userid": publicacion.UserID,
		"mensaje": publicacion.Mensaje,
		"fecha": publicacion.Fecha,
	}

	result, err := coleccion.InsertOne(ctx, registro)

	if err != nil {
		return "", false, err
	}

	objID, _ := result.InsertedID.(primitive.ObjectID)

	return objID.String(), true, nil

}
