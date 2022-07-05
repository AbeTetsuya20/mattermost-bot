package main

import (
	"fmt"
	"rat/domain"
)

func main() {
	fmt.Println("Hello World")

	user := domain.User{
		Email:  "a",
		Name:   "b",
		Column: "c",
	}
}
