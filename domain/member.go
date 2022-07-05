package domain

import "fmt"

// create user
type User struct {
	Email  string `yaml:"email"`
	Name   string `yaml:"name"`
	Column int    `yaml:"column"`
}

func NewUser(email string, name string, column int) *User {
	return &User{
		Email:  email,
		Name:   name,
		Column: column,
	}
}

type Users struct {
	Users []*User `yaml:"users"`
}

func PrintUsers(users *Users) {
	for _, user := range users.Users {
		fmt.Println(user)
	}
}
