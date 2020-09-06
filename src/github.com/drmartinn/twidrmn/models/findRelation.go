package models

/*FindRelation tiene true o false que se obtiene al consultar la relacion entre dos usuarios*/
type FindRelation struct {
	Status bool `json:"status"`
}
