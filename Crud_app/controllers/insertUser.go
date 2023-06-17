package controllers

import (
	"context"
	"fmt"
	"net/http"

	"Crud_app/configs"

	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

	fmt.Println("Username: ", Username, "Password: ", Password)

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
	_, err = collection.InsertOne(context.Background(), document)
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
