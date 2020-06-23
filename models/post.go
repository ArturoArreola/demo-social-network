package models

// Post es el model del body que llega en la petición
type Post struct {

	Mensaje string `bson:"mensaje" json:"mensaje"`

}
