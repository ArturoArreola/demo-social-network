package models

// Relacion es la estructura para grabar la relación de un usuario con otro
type Relacion struct {

	UsuarioID 			string `bson:"usuarioid", json: "usuarioId"`
	UsuarioRelacionID 	string `bson:"usuariorelacionid", json: "usuarioRelacionId"`

}