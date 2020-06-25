package routers

import (
	"errors"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/drmartinn/twidrmn/bd"
	"github.com/drmartinn/twidrmn/models"
)

/*Email valor del email usuado en todos los enpoints*/
var Email string

/*IDUsuario es el ID devuelto del modelo, que se usara en todos los enpoints*/
var IDUsuario string

/*ProcessToken funcion para analizar el estado de un token*/
func ProcessToken(token string) (*models.Claim, bool, string, error) {
	miClave := []byte("K3yG3n3rat3ByDrmartin3")
	claims := &models.Claim{}
	splitToken := strings.Split(token, "Bearer")
	if len(splitToken) != 2 {
		return claims, false, string(""), errors.New("formato de token invalido")
	}
	token = strings.TrimSpace(splitToken[1])
	tkn, err := jwt.ParseWithClaims(token, claims, func(tokenFunc *jwt.Token) (interface{}, error) {
		return miClave, nil
	})
	if err == nil {
		_, encontrado, _ := bd.CheckExistUser(claims.Email)
		if encontrado == true {
			Email = claims.Email
			IDUsuario = claims.ID.Hex()
		}
		return claims, encontrado, IDUsuario, nil
	}
	if !tkn.Valid {
		return claims, false, string(""), errors.New("token invalido")
	}
	return claims, false, string(""), err
}
