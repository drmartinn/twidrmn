package bd

import (
	"context"
	"log"

	"sync"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type connection struct {
	conn *mongo.Client
}

var (
	once     sync.Once
	instance *connection
	lock     = &sync.Mutex{}
)

/*MongoCN es el objeto de conexion a la BD */
var MongoCN = GetInstanceBD()
var clienteOptions = options.Client().ApplyURI("mongodb+srv://drmartinn:@cluster0-rtfqn.mongodb.net/test")

func init() {
	once.Do(
		func() {
			client, err := mongo.Connect(context.TODO(), clienteOptions)
			if err != nil {
				log.Fatal(err.Error())
				instance.conn = client
			}
			err = client.Ping(context.TODO(), nil)
			if err != nil {
				log.Fatal(err.Error())
				instance.conn = client
			}
			log.Printf("Conexión Exitosa con la BD\n")
			instance = &connection{conn: client}
		})
}

/*GetInstanceBD es una funcion para conectarse a la BD*/
func GetInstanceBD() *connection {
	return instance
}

/*CheckConnection es una funcion para determinar el estado de la conexion*/
func CheckConnection(connection *connection) int {
	if connection != nil {
		lock.Lock()
		defer lock.Unlock()
		err := connection.conn.Ping(context.TODO(), nil)
		if err != nil {
			return 0
		}
		log.Println("Ping exitoso")
		return 1
	} else {
		return 0
	}
}
