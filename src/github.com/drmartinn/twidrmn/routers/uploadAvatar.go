package routers

import (
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/drmartinn/twidrmn/bd"
	"github.com/drmartinn/twidrmn/models"
)

/*UploadAvatar funcion apra subir imagen del avatar*/
func UploadAvatar(w http.ResponseWriter, r *http.Request) {
	file, handler, err := r.FormFile("avatar")
	var extension = strings.Split(handler.Filename, ".")[1]
	var archivo string = "uploads/avatars/" + IDUsuario + "." + extension

	f, err := os.OpenFile(archivo, os.O_RDONLY|os.O_CREATE, 0666)
	if err != nil {
		http.Error(w, "Error al subir la imagen ! "+err.Error(), http.StatusInternalServerError)
		return
	}
	_, err = io.Copy(f, file)
	if err != nil {
		http.Error(w, "Error copiando la imagen ! "+err.Error(), http.StatusInternalServerError)
		return
	}

	var usuario models.User
	var status bool
	usuario.Avatar = IDUsuario + "." + extension
	status, err = bd.UpdateRecord(usuario, IDUsuario)
	if err != nil || status == false {
		http.Error(w, "Error guardando el usuario en BD "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusAccepted)

}
