package bd

import (
	"context"
	"time"

	"github.com/santicuevas/TwitterGO/models"
	"go.mongodb.org/mongo-driver/bson"
)

func LeoTweetsFollowers(Id string, pagina int) ([]models.DevuelvoTweetsFollowers, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	db := MongoCN.Database("TwitterGO")
	col := db.Collection("relation")

	skip := (pagina - 1) * 20

	condiciones := make([]bson.M, 0)
	condiciones = append(condiciones, bson.M{"$match": bson.M{"userId": Id}})
	condiciones = append(condiciones, bson.M{
		"$lookup": bson.M{
			"from":         "tweet",
			"localField":   "userRelationId",
			"foreignField": "userId",
			"as":           "tweet",
		},
	})
	condiciones = append(condiciones, bson.M{"$unwind": "tweet"})
	condiciones = append(condiciones, bson.M{"$sort": bson.M{"date": -1}}) //de menor a mayor es = 1
	condiciones = append(condiciones, bson.M{"$skip": skip})
	condiciones = append(condiciones, bson.M{"$limit": 20})

	cursor, err := col.Aggregate(ctx, condiciones)
	var result []models.DevuelvoTweetsFollowers
	err = cursor.All(ctx, &result)
	if err != nil {
		return result, false
	}
	return result, true
}
