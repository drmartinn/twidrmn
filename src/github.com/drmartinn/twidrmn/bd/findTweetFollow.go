package bd

import (
	"context"
	"time"

	"github.com/drmartinn/twidrmn/models"
	"go.mongodb.org/mongo-driver/bson"
)

/*FindTweetFollow funcion para obtener todos los tweets de los que sigo*/
func FindTweetFollow(ID string, pagina int) ([]models.TweetsFollow, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	dbInstance := GetInstanceBD()
	db := dbInstance.conn.Database("twidrmn")
	col := db.Collection("relation")

	skip := (pagina - 1) * 20
	condiciones := make([]bson.M, 0)
	condiciones = append(condiciones, bson.M{"$match": bson.M{"usuarioid": ID}})
	condiciones = append(condiciones, bson.M{
		"$lookup": bson.M{
			"from":         "tweet",
			"localField":   "usuariorelationid",
			"foreignField": "userid",
			"as":           "tweet",
		}})
	condiciones = append(condiciones, bson.M{"$unwind": "$tweet"})
	condiciones = append(condiciones, bson.M{"$sort": bson.M{"tweet.date": -1}})
	condiciones = append(condiciones, bson.M{"$skip": skip})
	condiciones = append(condiciones, bson.M{"$limit": 20})
	cursor, err := col.Aggregate(ctx, condiciones)
	var result []models.TweetsFollow
	err = cursor.All(ctx, &result)
	if err != nil {
		return result, false
	}
	return result, true
}
