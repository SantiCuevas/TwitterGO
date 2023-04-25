package bd

import (
	"context"
	"time"

	"github.com/santicuevas/TwitterGO/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// InsertoRegistro es la parada final con la BD para insertar los datos del usuario
func InsertoRegistro(u models.User) (string, bool, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittergo")
	col := db.Collection("users")

	u.Password, _ = EncriptarPassword(u.Password)

	result, err := col.InsertOne(ctx, u)
	if err != nil {
		return "", false, err
	}

	ObjID, _ := result.InsertedID.(primitive.ObjectID)
	return ObjID.String(), true, nil
}
