package database

import (
	"context"
	"time"

	"github.com/drc288/go_mux/models"
	"go.mongodb.org/mongo-driver/bson"
)

/**
 * CheckEmail - Check if email are using in database
 * @email: the email to verify
 *
 * Return: (user, if exists, id of user)
 */
func CheckEmail(email string) (models.User, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database(nameDB)
	collection := db.Collection(collectionUsers)

	condition := bson.M{"email": email}

	var result models.User

	err := collection.FindOne(ctx, condition).Decode(&result)
	ID := result.ID.Hex()

	if err != nil {
		return result, false, ID
	}

	return result, true, ID
}
