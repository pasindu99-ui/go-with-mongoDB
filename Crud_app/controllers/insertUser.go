package controllers

import (
	"context"

	"Crud_app/configs"

	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InsertUser(c echo.Context) error {

	username := c.FormValue("user_name")
	email := c.FormValue("email")
	password := c.FormValue("password")

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
		username: username,
		email:    email,
		password: password,
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
