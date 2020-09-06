package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/drmartinn/twidrmn/bd"
)

/*FindTweetsFollow funcion para obtener los tweets de mis seguidores*/
func FindTweetsFollow(w http.ResponseWriter, r *http.Request) {
	if len(r.URL.Query().Get("pagina")) < 1 {
		http.Error(w, "Debe enviar el parametro página", http.StatusBadRequest)
		return
	}
	pagina, err := strconv.Atoi(r.URL.Query().Get("pagina"))
	if err != nil {
		http.Error(w, "Debe enviar el parametro página como entero mayor a 0", http.StatusBadRequest)
		return
	}
	respuesta, correcto := bd.FindTweetFollow(IDUsuario, pagina)
	if correcto == false {
		http.Error(w, "Error al leer los tweets", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Contect-type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(respuesta)
}
