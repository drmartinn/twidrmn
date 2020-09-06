package bd

import (
	"context"
	"fmt"
	"time"

	"github.com/drmartinn/twidrmn/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*FindUsers funcion para obtener listado de usuario*/
func FindUsers(ID string, page int64, search string, tipo string) ([]*models.User, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	dbInstance := GetInstanceBD()
	db := dbInstance.conn.Database("twidrmn")
	col := db.Collection("users")

	var resultado []*models.User
	findOption := options.Find()
	findOption.SetSkip((page - 1) * 20)
	findOption.SetLimit(20)

	query := bson.M{
		"name": bson.M{"$regex": `(?i)` + search},
	}
	cur, err := col.Find(ctx, query, findOption)
	if err != nil {
		fmt.Println(err.Error())
		return resultado, false
	}
	var encontrado, incluir bool
	for cur.Next(ctx) {
		var s models.User
		err := cur.Decode(&s)
		if err != nil {
			fmt.Println(err.Error())
			return resultado, false
		}
		var r models.Relation
		r.UsuarioID = ID
		r.UsuarioRelationID = s.ID.Hex()

		incluir = false
		encontrado, err = GetRelation(r)
		if tipo == "new" && encontrado == false {
			incluir = true
		}

		if tipo == "follow" && encontrado == true {
			incluir = true
		}

		if r.UsuarioRelationID == ID {
			incluir = false
		}

		if incluir == true {
			s.Password = ""
			s.Biography = ""
			s.PageWeb = ""
			s.Location = ""
			s.Banner = ""
			s.Email = ""
			resultado = append(resultado, &s)
		}
	}

	err = cur.Err()
	if err != nil {
		fmt.Println(err.Error())
		return resultado, false
	}

	cur.Close(ctx)
	return resultado, true

}
