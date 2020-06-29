package routers

import (
	"encoding/json"
	"net/http"

	"github.com/drmartinn/twidrmn/bd"
	"github.com/drmartinn/twidrmn/models"
)

/*ModProfile funcion para modificar el perfil*/
func ModProfile(w http.ResponseWriter, r *http.Request) {
	var t models.User
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Datos incorrectos "+err.Error(), http.StatusBadGateway)
		return
	}
	var status bool
	status, err = bd.UpdateRecord(t, IDUsuario)
	if err != nil {
		http.Error(w, "Ocurrio un error al intentar modificar el registro, reintente nuevamente "+err.Error(), 400)
		return
	}
	if status == false {
		http.Error(w, "No se ha logrado modificar el registro del usuario", http.StatusBadGateway)
		return
	}
	w.WriteHeader(http.StatusCreated)

}
