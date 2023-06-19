package controllers

import (
	"context"
	"net/http"

	"Crud_app/configs"

	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/bson"
)

func InsertUser(c echo.Context) error {

	var requestData User

	// Bind the JSON data from the request body into the requestData variable
	if err := c.Bind(&requestData); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON data")
	}

	// Access the values from the requestData variable
	Username := requestData.Username
	email := requestData.Email
	Password := requestData.Password

	client := configs.ConnectDb()

	// Access the collection
	collection := client.Database("Users").Collection("User_details")

	// Create a new document
	person := person{
		username: Username,
		email:    email,
		password: Password,
	}

	document := bson.M{
		"username": person.username,
		"email":    person.email,
		"password": person.password,
	}

	// Insert the document
	_, err := collection.InsertOne(context.Background(), document)
	if err != nil {
		return err
	}

	return c.JSON(200, "Inserted a single document: "+person.username)
}

type person struct {
	username string
	email    string
	password string
}
