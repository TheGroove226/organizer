package handlers

import (
	"fmt"
	"log"
	"organizer/pkg/event"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

// PostEvent Handler
func PostEvent(c echo.Context) error {

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["id"].(string)
	fmt.Println("Welcome " + name + "!")

	e := event.NewEvent()

	if err := c.Bind(e); err != nil {
		log.Printf("%s", err)
		return err
	}

	eid, err := e.CreateEvent(claims["id"])

	if err != nil {
		log.Fatal(err)
	}

	return c.JSON(201, eid)
}
