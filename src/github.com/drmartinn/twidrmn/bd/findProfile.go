package bd

import (
	"context"
	"fmt"
	"time"

	"github.com/drmartinn/twidrmn/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*FindProfile busca un perfil del usuario*/
func FindProfile(ID string) (models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()
	dbInstance := GetInstanceBD()
	db := dbInstance.conn.Database("twidrmn")
	col := db.Collection("users")
	var perfil models.User
	objID, _ := primitive.ObjectIDFromHex(ID)
	condicion := bson.M{
		"_id": objID,
	}
	err := col.FindOne(ctx, condicion).Decode(&perfil)
	perfil.Password = ""
	if err != nil {
		fmt.Println("Registro no encontrado " + err.Error())
		return perfil, err
	}
	return perfil, nil
}
