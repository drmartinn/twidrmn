package routers

import (
	"io"
	"net/http"
	"os"

	"github.com/drmartinn/twidrmn/bd"
)

/*GetAvatar envia el avatar al http*/
func GetAvatar(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parametro ID ", http.StatusBadRequest)
		return
	}
	perfil, err := bd.FindProfile(ID)
	if err != nil {
		http.Error(w, "Usuario no encontrado", http.StatusInternalServerError)
		return
	}
	OpenFile, err := os.Open("uploads/avatars/" + perfil.Avatar)
	if err != nil {
		http.Error(w, "Imagen no encontrada", http.StatusInternalServerError)
		return
	}
	_, err = io.Copy(w, OpenFile)
	if err != nil {
		http.Error(w, "Error al copiar la imagen", http.StatusInternalServerError)
		return
	}
}
