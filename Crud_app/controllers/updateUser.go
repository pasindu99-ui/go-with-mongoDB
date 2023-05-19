package controllers

import (
	"Crud_app/configs"
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func UpdateUser() {
	url := configs.EnvMongoURI()
	// Set up MongoDB connection
	clientOptions := options.Client().ApplyURI(url)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	// Access the database and collection
	database := client.Database("<database>")
	collection := database.Collection("<collection>")

	// Define the filter to identify the document(s) to update
	filter := bson.M{"name": "John Doe"}

	// Define the update operation
	update := bson.M{
		"$set": bson.M{
			"email": "newemail@example.com",
		},
	}

	// Perform the update operation
	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Matched %d document(s) and modified %d document(s)\n", result.MatchedCount, result.ModifiedCount)

}
