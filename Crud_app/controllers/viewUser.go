package controllers

import (
	"Crud_app/configs"
	"context"
	"fmt"

	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ViewUser(c echo.Context) error {
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

	// Read data from the collection
	cursor, err := AccessToDb.Find(context.Background(), bson.M{})
	if err != nil {
		return err
	}
	defer cursor.Close(context.Background())

	// Iterate over the cursor to access the documents
	for cursor.Next(context.Background()) {
		var result bson.M
		if err := cursor.Decode(&result); err != nil {
			return err
		}
		fmt.Println(result)
		return c.JSON(200, result)
	}

	if err := cursor.Err(); err != nil {
		return err
	}
	return c.JSON(200, "Viewed all documents: ")
}
