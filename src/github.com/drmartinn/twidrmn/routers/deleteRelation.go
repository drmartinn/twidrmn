package routers

import (
	"net/http"

	"github.com/drmartinn/twidrmn/bd"
	"github.com/drmartinn/twidrmn/models"
)

/*DeleteRelation Funcion para eliminar una relacion entre usuarios*/
func DeleteRelation(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "El id es obligatorio", http.StatusBadRequest)
		return
	}
	var t models.Relation
	t.UsuarioID = IDUsuario
	t.UsuarioRelationID = ID
	status, err := bd.DeleteRelation(t)
	if err != nil {
		http.Error(w, "Error borrando relacion "+err.Error(), http.StatusInternalServerError)
		return
	}
	if status == false {
		http.Error(w, "Error borrando relacion ", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
