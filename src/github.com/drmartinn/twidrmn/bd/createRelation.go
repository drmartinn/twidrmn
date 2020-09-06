package bd

import (
	"context"
	"time"

	"github.com/drmartinn/twidrmn/models"
)

/*CreateRelation funcion para crear una relacion entre usuarios*/
func CreateRelation(t models.Relation) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()
	dbInstance := GetInstanceBD()
	db := dbInstance.conn.Database("twidrmn")
	col := db.Collection("relation")

	_, err := col.InsertOne(ctx, t)
	if err != nil {
		return false, err
	}
	return true, nil
}
