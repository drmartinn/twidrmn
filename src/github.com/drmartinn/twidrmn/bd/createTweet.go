package bd

import (
	"context"
	"time"

	"github.com/drmartinn/twidrmn/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*CreateTweet funcion para crear un tweet*/
func CreateTweet(t models.Tweet) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()
	dbInstance := GetInstanceBD()
	db := dbInstance.conn.Database("twidrmn")
	col := db.Collection("tweet")

	registro := bson.M{
		"userID":  t.UserId,
		"message": t.Message,
		"date":    t.Date,
	}
	result, err := col.InsertOne(ctx, registro)
	if err != nil {
		return "", false, err
	}
	objID, _ := result.InsertedID.(primitive.ObjectID)
	return objID.String(), true, nil
}
