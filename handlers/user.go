package handlers

import (
	"organizer/pkg/user"

	"github.com/labstack/echo"
)

// PostUser Func
func PostUser(c echo.Context) error {
	u := user.NewUser()

	if err := c.Bind(u); err != nil {
		return err
	}

	uid, err := u.CreateUser()

	if err != nil {
		return err
	}

	u.ID = uid

	return c.JSON(201, u)
}
