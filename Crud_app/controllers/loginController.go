package controllers

import (
	"Crud_app/configs"
	"context"
	"net/http"

	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/bson"
)

type User struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func LoginControll(c echo.Context) error {

	var requestData User

	// Bind the JSON data from the request body into the requestData variable
	if err := c.Bind(&requestData); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON data")
	}

	// Access the values from the requestData variable
	Username := requestData.Username
	Password := requestData.Password

	if Username == "" {
		// will be printed on the console, since str1 is empty
		return c.JSON(400, "user-name is empty")
	} else if Password == "" {
		// will be printed on the console, since str1 is empty
		return c.JSON(400, "user-password is empty")
	} else {

		filter := bson.M{"username": requestData.Username}

		client := configs.ConnectDb()

		// Access the database and collection
		AccessToDb := client.Database("Users").Collection("User_details")
		var users []User
		cursor, err := AccessToDb.Find(context.Background(), filter)
		if err != nil {
			return err
		}
		if err := cursor.All(context.Background(), &users); err != nil {
			return err
		}

		if len(users) == 0 {
			return c.JSON(400, "User not found")

		} else {
			checkPassword := users[0].Password
			println(users[0].Password)

			if checkPassword == Password {
				return c.JSON(200, "Login Success")
			} else {
				return c.JSON(400, "Login Failed")
			}
		}

	}
}
