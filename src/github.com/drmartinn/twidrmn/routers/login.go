package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/drmartinn/twidrmn/bd"
	"github.com/drmartinn/twidrmn/jwt"
	"github.com/drmartinn/twidrmn/models"
)

/*Login router que realiza el login*/
func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "User or password are invalid "+err.Error(), 400)
		return
	}
	if len(user.Email) == 0 {
		http.Error(w, "Field email is required", 400)
		return
	}
	document, find := bd.AttemptLogin(user.Email, user.Password)
	if find == false {
		http.Error(w, "User or password are invalid", 400)
		return
	}
	jwtKey, err := jwt.GenerateJWT(document)
	if err != nil {
		http.Error(w, "Error on generate JWT "+err.Error(), 400)
		return
	}
	response := models.ResponseLogin{
		Token: jwtKey,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)

	expirationTime := time.Now().Add(24 * time.Hour)
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   jwtKey,
		Expires: expirationTime,
	})

}
