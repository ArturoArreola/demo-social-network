package models

import "time"

// GrabarPost es la estructura que va a tener Twee en la base de datos
type GrabarPost struct {

	UserID 	string 		`bson:"userid" json:"userid,omitempty"`
	Mensaje string		`bson:"mensaje" json:"mensaje,omitempty"`
	Fecha 	time.Time	`bson:"fecha" json:"fecha,omitempty"`

}
