package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User model
type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name      string             `bson:"name" json:"name, omitempty"`
	LastName  string             `bson:"lastName" json:"lastName, omitempty"`
	BirthDate time.Time          `bson:"birthDate" json:"birthDate, omitempty"`
	Email     string             `bson:"email" json:"email"`
	Password  string             `bson:"password" json:"password, omitempty"`
	Avatar    string             `bson:"avatar" json:"avatar, omitempty"`
	Banner    string             `bson:"banner" json:"banner, omitempty"`
	Biography string             `bson:"biografia" json:"biography, omitempty"`
	Country   string             `bson:"country" json:"country, omitempty"`
	WebSite   string             `bson:"website" json:"webSite, omitempty"`
}
