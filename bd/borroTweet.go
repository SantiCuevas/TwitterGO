package bd

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// BorroTweet borra un tweet
func BorroTweet(Id string, userId string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()
	db := MongoCN.Database("TwitterGO")
	col := db.Collection("tweet")

	objId, _ := primitive.ObjectIDFromHex(Id)

	condicion := bson.M{
		"_id":    objId,
		"userId": userId,
	}
	_, err := col.DeleteOne(ctx, condicion)
	return err
}
