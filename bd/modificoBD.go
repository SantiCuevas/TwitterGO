package bd

import (
	"context"
	"time"

	"github.com/santicuevas/TwitterGO/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ModificoRegistro me permite modificar el perfil del usuario
func ModificoRegistro(u models.User, ID string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := MongoCN.Database("TwitterGO")
	col := db.Collection("users")

	registro := make(map[string]interface{})
	if len(u.Name) > 0 {
		registro["name"] = u.Name
	}
	if len(u.SurName) > 0 {
		registro["surName"] = u.SurName
	}
	registro["birthDate"] = u.BirthDate
	if len(u.Banner) > 0 {
		registro["banner"] = u.Banner
	}
	if len(u.Location) > 0 {
		registro["location"] = u.Location
	}
	if len(u.Avatar) > 0 {
		registro["avatar"] = u.Avatar
	}
	if len(u.Biography) > 0 {
		registro["biography"] = u.Biography
	}
	if len(u.WebSite) > 0 {
		registro["webSite"] = u.WebSite
	}

	updtString := bson.M{
		"$set": registro,
	}

	objID, _ := primitive.ObjectIDFromHex(ID)
	filtro := bson.M{"_id": bson.M{"$eq": objID}}

	_, err := col.UpdateOne(ctx, filtro, updtString)
	if err != nil {
		return false, err
	}
	return true, nil
}
