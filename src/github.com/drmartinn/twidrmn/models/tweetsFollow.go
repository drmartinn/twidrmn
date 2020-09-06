package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*TweetsFollow estructura para almacenar listado de tweets*/
type TweetsFollow struct {
	ID                primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	Usuario           string             `bson:"usuarioid" json:"userId,omitempty"`
	UsuarioRelationID string             `bson:"usuariorelationid" json:"userRelationId,omitempty"`
	Tweet             struct {
		Mensaje string    `bson:"message" json:"message,omitempty"`
		Date    time.Time `bson:"date,omitempty" json:"date"`
		ID      string    `bson:"_id" json:"_id,omitempty"`
	}
}
