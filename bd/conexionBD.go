package bd

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoCN es la variable para realizar la conexión a la función ConectarDB
var MongoCN = ConectarBD()
var clientOptions = options.Client().ApplyURI("mongodb+srv://mongoUser:mongodb1310@demo-social-network-ljiln.mongodb.net/test")

// ConectarBD es la función para poder realizar la conexión a la base de datos
func ConectarBD() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	log.Printf("Conexión exitosa a MongoDB")
	return client
}

// CheckConnection es la función para dar pinga la base de datos
func CheckConnection() int {
	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}
	return 1
}
