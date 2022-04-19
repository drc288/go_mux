package database

import (
	"context"
	"time"

	"github.com/drc288/go_mux/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ModifyRegister(u models.User, ID string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database(nameDB)
	collection := db.Collection(collectionUsers)

	register := make(map[string]interface{})

	register["birthDate"] = u.BirthDate
	if len(u.Name) > 0 {
		register["name"] = u.Name
	}

	if len(u.LastName) > 0 {
		register["lastName"] = u.LastName
	}

	if len(u.Avatar) > 0 {
		register["avatar"] = u.Avatar
	}

	if len(u.Country) > 0 {
		register["country"] = u.Country
	}

	if len(u.Biography) > 0 {
		register["biography"] = u.Biography
	}

	if len(u.Banner) > 0 {
		register["banner"] = u.Banner
	}

	if len(u.WebSite) > 0 {
		register["webSite"] = u.WebSite
	}

	updateRegister := bson.M{
		"$set": register,
	}

	objID, _ := primitive.ObjectIDFromHex(ID)
	filter := bson.M{
		"_id": bson.M{
			// Filter _id == ID
			"$eq": objID,
		},
	}

	_, err := collection.UpdateOne(ctx, filter, updateRegister)
	if err != nil {
		return false, err
	}

	return true, nil
}
