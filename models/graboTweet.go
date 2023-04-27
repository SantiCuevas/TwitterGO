package models

import "time"

//GraboTweet es el formato que tienen los tweets en la BD
type GraboTweet struct {
	UserId  string    `bson:"userId" json:"userId,omitempty"`
	Message string    `bson:"message" json:"message,omitempty"`
	Date    time.Time `bson:"date" json:"date,omitempty"`
}
