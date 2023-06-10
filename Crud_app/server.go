package main

import (
	"Crud_app/configs"
	"Crud_app/routes"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	configs.ConnectDb()

	routes.UserRoute(e)
	e.Logger.Fatal(e.Start(":6000"))
}
