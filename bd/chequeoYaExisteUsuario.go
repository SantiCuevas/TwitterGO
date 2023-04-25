package bd

import (
	"context"
	"time"

	"github.com/santicuevas/TwitterGO/models"
	"go.mongodb.org/mongo-driver/bson"
)

// ChequeoYaExisteUsuario recibe un mail de parametro y chequea si existe en la BD
func ChequeoYaExisteUsuario(email string) (models.User, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("TwitterGO")
	col := db.Collection("users")

	condicion := bson.M{"email": email}

	var resultado models.User

	err := col.FindOne(ctx, condicion).Decode(&resultado)
	ID := resultado.ID.Hex()
	if err != nil {
		return resultado, false, ID
	}
	return resultado, true, ID
}
