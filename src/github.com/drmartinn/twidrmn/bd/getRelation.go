package bd

import (
	"context"
	"fmt"
	"time"

	"github.com/drmartinn/twidrmn/models"
	"go.mongodb.org/mongo-driver/bson"
)

/*GetRelation funcion para obtener mis relaciones*/
func GetRelation(t models.Relation) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()
	dbInstance := GetInstanceBD()
	db := dbInstance.conn.Database("twidrmn")
	col := db.Collection("relation")

	condicion := bson.M{
		"usuarioid":         t.UsuarioID,
		"usuariorelationid": t.UsuarioRelationID,
	}

	var resultado models.Relation
	fmt.Println(resultado)
	err := col.FindOne(ctx, condicion).Decode(&resultado)
	if err != nil {
		fmt.Println(err.Error())
		return false, err
	}
	return true, nil
}
