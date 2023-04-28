package bd

import (
	"context"
	"fmt"
	"time"

	"github.com/santicuevas/TwitterGO/models"
	"go.mongodb.org/mongo-driver/bson"
)

// Consulta la relacion entre 2 usuarios
func ConsultoRelacion(t models.Relation) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	db := MongoCN.Database("TwitterGO")
	col := db.Collection("relation")

	condicion := bson.M{
		"userId":         t.UserId,
		"userRelationId": t.UserRelationId,
	}

	var resultado models.Relation
	fmt.Println(resultado)
	err := col.FindOne(ctx, condicion).Decode(&resultado)
	if err != nil {
		fmt.Println(err.Error())
		return false, err
	}

	return true, nil
}
