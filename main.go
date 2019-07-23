package main

import (
	"organizer/handlers"

	"github.com/labstack/echo"
)

func main() {

	api := echo.New()

	api.POST("/signup", handlers.PostUser)

	api.Start(":3000")
}
