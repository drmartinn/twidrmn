package main

import (
	"log"

	"github.com/drmartinn/twidrmn/bd"
	"github.com/drmartinn/twidrmn/handlers"
)

func main() {
	connectionBD := bd.GetInstanceBD()
	if bd.CheckConnection(connectionBD) == 0 {
		log.Fatal("No hay conexión a la BD")
		return
	}

	handlers.Manejadores()
}
