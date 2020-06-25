package bd

import (
	"context"
	"fmt"
	"time"

	"github.com/ArturoArreola/demo-social-network/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// LeerUsuariosTodos lee los usuarios registrados en la base de datos, hace el filtro  para retornar
//		unicamente los que están relacionados con el usuario que lanza la petición
func LeerUsuariosTodos(ID string, page int64, search string, tipo string) ([]*models.Usuario, bool) {

	ctx, cancel := context.WithTimeout(context.Background(),time.Second*15)
	defer cancel()

	db := MongoCN.Database("social-network-db")
	coleccion := db.Collection("usuarios")

	var resultados []*models.Usuario
	findoptions := options.Find()
	findoptions.SetSkip((page-1)*20)
	findoptions.SetLimit(20)

	query := bson.M{
		"nombre": bson.M{"$regex": `(?i)` + search},
	}

	cursor, err := coleccion.Find(ctx, query, findoptions)
	if err != nil {
		fmt.Println(err.Error())
		return resultados, false
	}

	var encontrado, incluir bool

	for cursor.Next(ctx){
		var usuario models.Usuario
		err := cursor.Decode(&usuario)
		if err != nil {
			fmt.Println(err.Error())
			return resultados, false
		}

		var r models.Relacion
		r.UsuarioID = ID
		r.UsuarioRelacionID = usuario.ID.Hex()

		incluir = false

		encontrado, err = ConsultaRelacion(r)
		if tipo == "new" && encontrado == false {
			incluir = true
		}

		if tipo == "follow" && encontrado == true {
			incluir = true
		}

		if r.UsuarioRelacionID == ID {
			incluir = false
		}

		if incluir == true {
			usuario.Password 	= ""
			usuario.Biografia 	= ""
			usuario.SitioWeb 	= ""
			usuario.Ubicacion 	= ""
			usuario.Banner 		= ""
			usuario.Email 		= ""

			resultados = append(resultados, & usuario)
		}
	}

	err = cursor.Err()
	if err != nil {
		fmt.Println(err.Error())
		return resultados, false
	}

	return resultados, true

}