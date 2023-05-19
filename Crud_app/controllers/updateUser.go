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

func UpdateUser(c echo.Context) error {
	user_name := c.FormValue("user_name")
	email := c.FormValue("email")

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
	AccessToDb := client.Database("Users").Collection("User_details")

	// Define the filter to identify the document(s) to update
	filter := bson.M{"user_name": user_name}

	// Define the update operation
	update := bson.M{
		"$set": bson.M{
			"email": email,
		},
	}

	// Perform the update operation
	result, err := AccessToDb.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}

	log.Printf("Matched %d document(s) and modified %d document(s)\n", result.MatchedCount, result.ModifiedCount)
	return c.JSON(200, "Updated a single document: ")

}
