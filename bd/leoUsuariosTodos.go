package bd

import (
	"context"
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

	var results []*models.Usuario

	findOptions := options.Find()
	findOptions.SetSkip((page - 1) * 20)
	findOptions.SetLimit(20)

	query := bson.M{
		"nombre": bson.M{"$regex": `(?i)` + search},
	}

	cur, err := coleccion.Find(ctx, query, findOptions)
	if err != nil {
		return results, false
	}

	var encontrado, incluir bool

	for cur.Next(ctx) {
		var s models.Usuario
		err := cur.Decode(&s)
		if err != nil {
			return results, false
		}

		var r models.Relacion
		r.UsuarioID = ID
		r.UsuarioRelacionID = s.ID.Hex()

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
			s.Password = ""
			s.Biografia = ""
			s.SitioWeb = ""
			s.Ubicacion = ""
			s.Banner = ""
			s.Email = ""

			results = append(results, &s)
		}
	}

	err = cur.Err()
	if err != nil {
		return results, false
	}
	cur.Close(ctx)
	return results, true

}