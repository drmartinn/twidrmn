package middlew

import (
	"net/http"

	"github.com/drmartinn/twidrmn/bd"
)

/*CheckBD funcion middlew que permite conocer el estado de la BD*/
func CheckBD(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		connection := bd.GetInstanceBD()
		if bd.CheckConnection(connection) == 0 {
			http.Error(w, "Failed to connect to BD", 500)
			return
		}
		next.ServeHTTP(w, r)
	}
}
