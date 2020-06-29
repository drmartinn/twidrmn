package models

import "time"

/*Tweet estructura para los tweets*/
type Tweet struct {
	UserId  string    `bson:"userid,omitempty" json:"userid"`
	Message string    `bson:"message,omitempty" json:"message"`
	Date    time.Time `bson:"date,omitempty" json:"date"`
}
