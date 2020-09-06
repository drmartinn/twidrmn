package routers

import (
	"net/http"

	"github.com/drmartinn/twidrmn/bd"
	"github.com/drmartinn/twidrmn/models"
)

/*CreateRelation realiza la creacion de una relacion entre usuario*/
func CreateRelation(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Debe ingresar id ", http.StatusBadRequest)
		return
	}
	var t models.Relation
	t.UsuarioID = IDUsuario
	t.UsuarioRelationID = ID
	status, err := bd.CreateRelation(t)
	if err != nil {
		http.Error(w, "Error creando relacion "+err.Error(), http.StatusInternalServerError)
		return
	}
	if status == false {
		http.Error(w, "Error creando relacion ", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
