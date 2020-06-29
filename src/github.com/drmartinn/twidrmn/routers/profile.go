package routers

import (
	"encoding/json"
	"net/http"

	"github.com/drmartinn/twidrmn/bd"
)

/*Profile permite extraer los valores del perfil*/
func Profile(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parámetro ID", http.StatusBadRequest)
		return
	}
	perfil, err := bd.FindProfile(ID)
	if err != nil {
		http.Error(w, "Ocurrio un error al intentar buscar el registro "+err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("context-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(perfil)
}
