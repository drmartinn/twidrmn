package models

/*TweetRender captura del body, el mensaje que llega*/
type TweetRender struct {
	Message string `bson:"message,omitempty" json:"Message"`
}
