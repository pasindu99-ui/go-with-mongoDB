package controllers

import (
	"Crud_app/configs"
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/bson"
)

func UpdateUser(c echo.Context) error {
	var requestData User

	// Bind the JSON data from the request body into the requestData variable
	if err := c.Bind(&requestData); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON data")
	}

	// Access the values from the requestData variable
	user_name := requestData.Username
	email := requestData.Email
	fmt.Println("Username: ", user_name, "email: ", email)
	client := configs.ConnectDb()

	// Access the database and collection
	AccessToDb := client.Database("Users").Collection("User_details")

	// Define the filter to identify the document(s) to update
	filter := bson.M{"username": user_name}

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
