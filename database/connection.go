package database

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var clientOptions = options.Client().ApplyURI("mongodb+srv://drose:a9UsarArLY5R6LA@gomux.t13uw.mongodb.net/test")
var MongoCN = MongoDBConnection()

var (
	nameDB          string = "twitter"
	collectionUsers string = "users"
)

// Create db connection
func MongoDBConnection() *mongo.Client {
	// Validate client conection
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	// Validate if the conection is up
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	log.Println("Connection succesfull in mongodb")
	return client
}

// Check db connection using ping
func CheckConnection() int {
	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}
	return 1
}
