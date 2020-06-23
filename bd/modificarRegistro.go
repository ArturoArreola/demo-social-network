package bd

import (
	"context"
	"time"

	"github.com/ArturoArreola/demo-social-network/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ModificarRegistro permite modificar el perfil del usuario
func ModificarRegistro (usuario models.Usuario, ID string) (bool, error){

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("social-network-db")
	coleccion := db.Collection("usuarios")

	registro := make(map[string]interface{})

	if len(usuario.Nombre) > 0 {
		registro["nombre"] = usuario.Nombre
	}

	if len(usuario.Apellidos) > 0 {
		registro["apellidos"] = usuario.Apellidos
	}

	if len(usuario.Avatar) > 0 {
		registro["avatar"] = usuario.Avatar
	}

	if len(usuario.Banner) > 0 {
		registro["banner"] = usuario.Banner
	}

	if len(usuario.Biografia) > 0 {
		registro["biografia"] = usuario.Biografia
	}

	if len(usuario.Ubicacion) > 0 {
		registro["ubicacion"] = usuario.Ubicacion
	}

	if len(usuario.SitioWeb) > 0 {
		registro["sitioWeb"] = usuario.SitioWeb
	}

	registro["fechaNacimiento"] = usuario.FechaNacimiento

	cadenaUpdate := bson.M{"$set":registro}

	objID, _ := primitive.ObjectIDFromHex(ID)

	filtro := bson.M{"_id":bson.M{"$eq":objID}}

	_, err := coleccion.UpdateOne(ctx, filtro, cadenaUpdate)

	if err != nil {
		return false, err
	}
	return true, nil
}