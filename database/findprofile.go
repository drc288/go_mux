package database

import (
	"context"
	"log"
	"time"

	"github.com/drc288/go_mux/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Find profile in database using the id
func FindProfile(ID string) (models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := MongoCN.Database(nameDB)
	collection := db.Collection(collectionUsers)

	var profile models.User

	objID, _ := primitive.ObjectIDFromHex(ID)

	condition := bson.M{
		"_id": objID,
	}

	err := collection.FindOne(ctx, condition).Decode(&profile)
	profile.Password = ""
	if err != nil {
		log.Panicln("Profile not finded")
		return profile, err
	}

	return profile, nil
}
