package bd

import (
	"context"
	"time"

	"github.com/drmartinn/twidrmn/models"
)

func DeleteRelation(t models.Relation) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()
	dbInstance := GetInstanceBD()
	db := dbInstance.conn.Database("twidrmn")
	col := db.Collection("relation")

	_, err := col.DeleteOne(ctx, t)
	if err != nil {
		return false, err
	}
	return true, nil
}
