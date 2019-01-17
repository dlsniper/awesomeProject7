package main

import (
	"fmt"

	"github.com/pkg/errors"
)

type User struct {
	ID        int
	Firstname string
	Lastname  string
	Job       string
}

func main() {
	u := User{
		ID:        1,
		Firstname: "Florin",
		Lastname:  "Patan",
		Job:       "Developer Advocate",
	}

	_ = errors.New("demo")

	fmt.Printf("User details %v\n", u)
}
