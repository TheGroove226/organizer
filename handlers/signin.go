package handlers

import (
	"fmt"
	"log"
	"net/http"
	"organizer/db"
	"organizer/pkg/user"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"

	"golang.org/x/crypto/bcrypt"
)

// PasswordCompare function
func PasswordCompare(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// Signin Handler
func Signin(c echo.Context) error {
	u := user.User{}
	checked := user.User{}
	err := c.Bind(&u)
	if err != nil {
		log.Printf("%s", err)
	}
	stmt := db.Get().QueryRow(checkUser, u.Email)
	if err != nil {
		log.Printf("%s", err)
	}
	err2 := stmt.Scan(&checked.ID, &checked.Email, &checked.Password)
	if err2 != nil {
		log.Printf("%s", err2)
	}

	if PasswordCompare(u.Password, checked.Password) {
		fmt.Printf("Welcome %s\n", checked.Email)

		token := jwt.New(jwt.SigningMethodHS256)
		claims := token.Claims.(jwt.MapClaims)
		claims["id"] = checked.ID
		claims["email"] = checked.Email
		claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

		t, err := token.SignedString([]byte("secret"))
		if err != nil {
			log.Printf("%s", err)
		}

		return c.JSON(http.StatusOK, map[string]string{
			"token": t,
		})
	}
	return echo.ErrUnauthorized

}

const checkUser = `SELECT id, email, password FROM users where email = ?`
