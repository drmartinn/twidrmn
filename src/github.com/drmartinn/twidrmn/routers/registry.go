package routers

import (
	"encoding/json"
	"net/http"

	"github.com/drmartinn/twidrmn/bd"
	"github.com/drmartinn/twidrmn/models"
)

/*Registry es la funcion para crear en la BD el registro de usuarios*/
func Registry(w http.ResponseWriter, r *http.Request) {
	var userToCreate models.User
	err := json.NewDecoder(r.Body).Decode(&userToCreate)
	if err != nil {
		http.Error(w, "Error on data recive "+err.Error(), 400)
		return
	}
	if len(userToCreate.Email) == 0 {
		http.Error(w, "Field emails is required ", 400)
		return
	}

	if len(userToCreate.Password) < 6 {
		http.Error(w, "The password will be more than 6 character ", 400)
		return
	}

	_, find, _ := bd.CheckExistUser(userToCreate.Email)
	if find == true {
		http.Error(w, "The email already exist ", 400)
		return
	}

	_, status, err := bd.RegistryRecord(userToCreate)
	if err != nil {
		http.Error(w, "Error on create user "+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "Registry user are unsuccessful "+err.Error(), 400)
		return
	}

	w.WriteHeader(http.StatusCreated)

}
