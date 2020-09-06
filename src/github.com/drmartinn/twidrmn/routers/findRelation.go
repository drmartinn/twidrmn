package routers

import (
	"encoding/json"
	"net/http"

	"github.com/drmartinn/twidrmn/bd"
	"github.com/drmartinn/twidrmn/models"
)

/*FindRelation funcion para buscar una relacion entre dos usuarios*/
func FindRelation(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	var t models.Relation
	t.UsuarioID = IDUsuario
	t.UsuarioRelationID = ID

	var respuesta models.FindRelation
	status, err := bd.GetRelation(t)
	if err != nil || status == false {
		respuesta.Status = false
	} else {
		respuesta.Status = true
	}
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(respuesta)

}
