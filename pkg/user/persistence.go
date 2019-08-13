package user

import (
	"fmt"
	"log"
	"organizer/db"

	"golang.org/x/crypto/bcrypt"
)

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// CreateUser func
func (u *User) CreateUser() (int, error) {
	fmt.Println("Executing while creating user")
	stmt, err := db.Get().Prepare(createUser)
	if err != nil {
		return 0, err
	}

	hash, err := hashPassword(u.Password)

	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()
	res, err := stmt.Exec(u.Email, hash)

	if err != nil {
		return 0, err
	}

	r, err := res.LastInsertId()

	if err != nil {
		return 0, err
	}

	return int(r), nil

}

const createUser = `INSERT INTO users (email, password) VALUES (?, ?)`
