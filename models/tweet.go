package models

//Tweet captura el body, el mensaje que nos llega
type Tweet struct {
	Message string `bson:"message" json:"message"`
}
