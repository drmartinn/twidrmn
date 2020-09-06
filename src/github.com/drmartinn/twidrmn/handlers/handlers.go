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
	router.HandleFunc("/findTweet", middlew.CheckBD(middlew.ValidateJwt(routers.FindTweet))).Methods("GET")
	router.HandleFunc("/deleteTweet", middlew.CheckBD(middlew.ValidateJwt(routers.DeleteTweet))).Methods("DELETE")

	router.HandleFunc("/uploadAvatar", middlew.CheckBD(middlew.ValidateJwt(routers.UploadAvatar))).Methods("POST")
	router.HandleFunc("/getAvatar", middlew.CheckBD(routers.GetAvatar)).Methods("GET")

	router.HandleFunc("/uploadBanner", middlew.CheckBD(middlew.ValidateJwt(routers.UploadBanner))).Methods("POST")
	router.HandleFunc("/getBanner", middlew.CheckBD(routers.GetBanner)).Methods("GET")

	router.HandleFunc("/createRelation", middlew.CheckBD(middlew.ValidateJwt(routers.CreateRelation))).Methods("POST")
	router.HandleFunc("/deleteRelation", middlew.CheckBD(middlew.ValidateJwt(routers.DeleteRelation))).Methods("DELETE")

	router.HandleFunc("/findRelation", middlew.CheckBD(middlew.ValidateJwt(routers.FindRelation))).Methods("GET")

	router.HandleFunc("/listUsers", middlew.CheckBD(middlew.ValidateJwt(routers.ListUsers))).Methods("GET")

	router.HandleFunc("/listTweetsFollow", middlew.CheckBD(middlew.ValidateJwt(routers.FindTweetsFollow))).Methods("GET")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	fmt.Println("Servidor iniciado")
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
