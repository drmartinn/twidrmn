package models

/*ResponseLogin estructura para almacenar el token respuesta del login*/
type ResponseLogin struct {
	Token string `json:"token,omitempty"`
}
