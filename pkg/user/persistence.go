package user

import "organizer/db"

// CreateUser func
func (u *User) CreateUser() (int, error) {
	stmt, err := db.Get().Prepare(createUser)
	if err != nil {
		return 0, err
	}

	defer stmt.Close()
	res, err := stmt.Exec(u.Email, u.Password)
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
