package controllers

import (
	"context"
	"log"

	"Crud_app/configs"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InsertUser() {
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

	// Access the collection
	collection := client.Database("Users").Collection("User_details")

	// Create a new document
	person := person{
		username: "John",
		email:    "pasindutt",
		password: "123",
	}

	document := bson.M{
		"name":     person.username,
		"email":    person.email,
		"password": person.password,
	}

	// Insert the document
	_, err = collection.InsertOne(context.Background(), document)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Data inserted successfully")
}

type person struct {
	username string
	email    string
	password string
}
