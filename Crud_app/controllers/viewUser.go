package controllers

import (
	"Crud_app/configs"
	"context"

	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/bson"
)

func ViewUser(c echo.Context) error {

	client := configs.ConnectDb()
	// Access the database and collection
	AccessToDb := client.Database("Users").Collection("User_details")

	// Read data from the collection
	cursor, err := AccessToDb.Find(context.Background(), bson.M{})
	if err != nil {
		return err
	}

	var results []bson.M

	defer cursor.Close(context.Background())

	// Iterate over the cursor to access the documents
	for cursor.Next(context.Background()) {
		var result bson.M
		if err := cursor.Decode(&result); err != nil {
			return err
		}
		results = append(results, result)

	}

	if err := cursor.Err(); err != nil {
		return err
	}
	return c.JSON(200, results)
}
