package controllers

import (
	"Crud_app/configs"
	"context"
	"log"

	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DeleteUser(c echo.Context) error {
	// Set up MongoDB connection
	url := configs.EnvMongoURI()
	// Set up MongoDB connection
	clientOptions := options.Client().ApplyURI(url)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return err
	}

	// Check the connection
	err = client.Ping(context.Background(), nil)
	if err != nil {
		return err
	}

	// Access the database and collection
	collection := client.Database("Users").Collection("User_details")

	// Define the filter to identify the document(s) to delete
	filter := bson.M{"user_name": "pasindu"}

	// Perform the delete operation
	result, err := collection.DeleteMany(context.Background(), filter)
	if err != nil {
		return err
	}

	log.Printf("Deleted %d document(s)\n", result.DeletedCount)
	return c.JSON(200, "Deleted a single document: ")

}
