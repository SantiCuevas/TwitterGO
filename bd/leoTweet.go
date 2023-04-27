package bd

import (
	"context"
	"log"
	"time"

	"github.com/santicuevas/TwitterGO/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func LeoTweets(ID string, pagina int64) ([]*models.DevuelvoTweets, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	db := MongoCN.Database("TwitterGO")
	col := db.Collection("tweet")

	var resultados []*models.DevuelvoTweets
	condicion := bson.M{
		"userId": ID,
	}

	opciones := options.Find()
	opciones.SetLimit(20)                              //Limite
	opciones.SetSort(bson.D{{Key: "date", Value: -1}}) //Orden
	opciones.SetSkip((pagina - 1) * 20)                //Salto

	cursor, err := col.Find(ctx, condicion, opciones)
	if err != nil {
		log.Fatal(err.Error())
		return resultados, false
	}

	for cursor.Next(context.TODO()) {
		var registro models.DevuelvoTweets
		err := cursor.Decode(&registro)
		if err != nil {
			return resultados, false
		}
		resultados = append(resultados, &registro)
	}
	return resultados, true
}
