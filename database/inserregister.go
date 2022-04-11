package database

import (
	"context"
	"time"

	"github.com/drc288/go_mux/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func InsertRegister(user models.User) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database(nameDB)
	collection := db.Collection(collectionUsers)

	user.Password, _ = EncryptPassowrd(user.Password)

	result, err := collection.InsertOne(ctx, user)
	if err != nil {
		return "", false, err
	}

	ObjID, _ := result.InsertedID.(primitive.ObjectID)
	return ObjID.String(), true, nil
}
