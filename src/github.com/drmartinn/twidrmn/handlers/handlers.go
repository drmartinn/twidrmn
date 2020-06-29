package handlers

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"

	"github.com/drmartinn/twidrmn/middlew"
	"github.com/drmartinn/twidrmn/routers"
)

/*Manejadores se asigna el puerto el handler, y se sube el servidor web*/
func Manejadores() {
	router := mux.NewRouter()

	router.HandleFunc("/registry", middlew.CheckBD(routers.Registry)).Methods("POST")
	router.HandleFunc("/login", middlew.CheckBD(routers.Login)).Methods("POST")
	router.HandleFunc("/profile", middlew.CheckBD(middlew.ValidateJwt(routers.Profile))).Methods("GET")
	router.HandleFunc("/updateProfile", middlew.CheckBD(middlew.ValidateJwt(routers.ModProfile))).Methods("PUT")
	router.HandleFunc("/tweet", middlew.CheckBD(middlew.ValidateJwt(routers.CreateTweet))).Methods("POST")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	fmt.Println("Servidor iniciado")
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
