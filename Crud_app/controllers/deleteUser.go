package controllers

import (
	"Crud_app/configs"
	"context"
	"log"
	"net/http"

	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/bson"
)

func DeleteUser(c echo.Context) error {
	var requestData User

	// Bind the JSON data from the request body into the requestData variable
	if err := c.Bind(&requestData); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON data")
	}

	// Access the values from the requestData variable
	user_name := requestData.Username

	client := configs.ConnectDb()
	// Access the database and collection
	collection := client.Database("Users").Collection("User_details")

	// Define the filter to identify the document(s) to delete
	filter := bson.M{"username": user_name}

	// Perform the delete operation
	result, err := collection.DeleteMany(context.Background(), filter)
	if err != nil {
		return err
	}

	log.Printf("Deleted %d document(s)\n", result.DeletedCount)
	return c.JSON(200, "Deleted a single document: ")

}
