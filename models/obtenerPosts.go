package models

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ObtenerPosts es la estructura para obtener los posts de un usuario
type ObtenerPosts struct {

	ID 		primitive.ObjectID 	`bson:"_id" json: "_id,omitempty"`
	UserID 	string 				`bson:"userid" json: "userId,omitempty"`
	Mensaje	string 				`bson:"mensaje" json: "mensaje,omitempty"`
	Fecha 	time.Time 			`bson:"fecha" json: "fecha,omitempty"`

}
