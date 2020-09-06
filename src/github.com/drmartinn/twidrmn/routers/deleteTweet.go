package routers

import (
	"net/http"

	"github.com/drmartinn/twidrmn/bd"
)

/*DeleteTweet permite borrar un tweet router*/
func DeleteTweet(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parametro ID", http.StatusBadRequest)
		return
	}
	err := bd.DeleteTweet(ID, IDUsuario)
	if err != nil {
		http.Error(w, "Error eliminando el tweet "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
