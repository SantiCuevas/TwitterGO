package bd

import (
	"context"
	"time"

	"github.com/santicuevas/TwitterGO/models"
)

// Graba la relacion en la BD
func InsertoRelacion(t models.Relation) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	db := MongoCN.Database("TwitterGO")
	col := db.Collection("relation")

	_, err := col.InsertOne(ctx, t)
	if err != nil {
		return false, err
	}

	return true, nil
}
