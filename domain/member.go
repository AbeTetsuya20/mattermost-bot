package domain

// create user
type User struct {
	Email  string `json:"email"`
	Name   string `json:"name"`
	Column string `json:"column"`
}

func NewUser(email string, name string, column string) *User {
	return &User{
		Email:  email,
		Name:   name,
		Column: column,
	}
}

type Users []*User
