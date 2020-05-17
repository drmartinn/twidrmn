package bd

import (
	"github.com/drmartinn/twidrmn/models"
	"golang.org/x/crypto/bcrypt"
)

/*AttemptLogin relaliza el chequeo  de login a la BD*/
func AttemptLogin(email string, password string) (models.User, bool) {
	usu, find, _ := CheckExistUser(email)
	if find == false {
		return usu, false
	}
	passwordBytes := []byte(password)
	passwordBD := []byte(usu.Password)
	err := bcrypt.CompareHashAndPassword(passwordBD, passwordBytes)
	if err != nil {
		return usu, false
	}
	return usu, true
}
