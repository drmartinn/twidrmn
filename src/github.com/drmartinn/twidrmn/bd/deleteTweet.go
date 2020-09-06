package bd

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*DeleteTweet funcion para eliminar un tweet*/
func DeleteTweet(ID string, UserID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()
	dbInstance := GetInstanceBD()
	db := dbInstance.conn.Database("twidrmn")
	col := db.Collection("tweet")

	objID, _ := primitive.ObjectIDFromHex(ID)
	condicion := bson.M{
		"_id":    objID,
		"userID": UserID,
	}
	_, err := col.DeleteOne(ctx, condicion)
	fmt.Println("Error eliminando twwt ID", ID)
	return err
}
