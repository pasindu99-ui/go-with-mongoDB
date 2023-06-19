package main

import (
	"Crud_app/routes"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	routes.UserRoute(e)
	e.Logger.Fatal(e.Start(":6000"))
}
