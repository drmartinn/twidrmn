package bd

import "golang.org/x/crypto/bcrypt"

func EncryptPassword(password string) (string, error) {
	costo := 8
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), costo)
	return string(bytes), err
}
