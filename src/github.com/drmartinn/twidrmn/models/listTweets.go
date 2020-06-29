package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*ListTweets listado de tweet*/
type ListTweets struct {
	ID      primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	UserID  string             `bson:"userid,omitempty" json:"userId"`
	Message string             `bson:"message,omitempty" json:"message"`
	Date    time.Time          `bson:"date,omitempty" json:"date"`
}
