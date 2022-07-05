package domain

// create user
type User struct {
	Email  string `json:"email"`
	Name   string `json:"name"`
	Column int    `json:"column"`
}

func NewUser(email string, name string, column int) *User {
	return &User{
		Email:  email,
		Name:   name,
		Column: column,
	}
}

type Users []*User
