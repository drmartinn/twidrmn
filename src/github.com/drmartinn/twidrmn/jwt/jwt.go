package jwt

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/drmartinn/twidrmn/models"
)

/*GenerateJWT funcion para generar los jwt*/
func GenerateJWT(user models.User) (string, error) {
	myKey := []byte("K3yG3n3rat3ByDrmartin3")
	payload := jwt.MapClaims{
		"email":     user.Email,
		"name":      user.Name,
		"lastName":  user.LastName,
		"birthDate": user.Birthdate,
		"_id":       user.ID.Hex(),
		"exp":       time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(myKey)
	if err != nil {
		return tokenStr, err
	}
	return tokenStr, nil
}
