package handlers

import (
	"log"
	"organizer/pkg/user"

	"github.com/labstack/echo"
)

// PostUser Func
func PostUser(c echo.Context) error {
	u := user.NewUser()

	if err := c.Bind(u); err != nil {
		log.Panic(err)
	}

	uid, err := u.CreateUser()

	if err != nil {
		log.Panic(err)
	}

	u.ID = uid

	return c.JSON(201, u)
}
