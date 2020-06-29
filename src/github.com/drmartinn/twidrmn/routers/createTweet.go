package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/drmartinn/twidrmn/bd"
	"github.com/drmartinn/twidrmn/models"
)

func CreateTweet(w http.ResponseWriter, r *http.Request) {
	var message models.TweetRender
	err := json.NewDecoder(r.Body).Decode(&message)
	if err != nil {
		http.Error(w, "Datos incorrectos "+err.Error(), http.StatusBadGateway)
		return
	}
	registro := models.Tweet{
		UserId:  IDUsuario,
		Message: message.Message,
		Date:    time.Now(),
	}
	_, status, err := bd.CreateTweet(registro)
	if err != nil {
		http.Error(w, "Ocurrio un error al intentar registrar el tweet "+err.Error(), 400)
		return
	}
	if status == false {
		http.Error(w, "No se ha logrado insertar el tweet ", 400)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
