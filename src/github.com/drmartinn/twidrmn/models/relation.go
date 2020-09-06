package models

/*Relation modelo para grabar la relacion de un usuario con otro*/
type Relation struct {
	UsuarioID         string `bson:"usuarioid" json:"usuarioId"`
	UsuarioRelationID string `bson:"usuariorelationid" json:"usuarioRelationId"`
}
