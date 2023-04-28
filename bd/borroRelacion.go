package bd

import (
	"context"
	"time"

	"github.com/santicuevas/TwitterGO/models"
)

// BorroRelacion borra la relacion de la BD
func BorroRelacion(t models.Relation) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	db := MongoCN.Database("TwitterGO")
	col := db.Collection("relation")

	_, err := col.DeleteOne(ctx, t)
	if err != nil {
		return false, err
	}

	return true, nil
}
