package user

// User sturcture
type User struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// NewUser Func
func NewUser() *User {
	u := &User{}
	return u
}
