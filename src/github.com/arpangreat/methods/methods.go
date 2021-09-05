package main

import (
	"fmt"
)

type User struct {
	Firstname, Lastname string
}

func (u User) Greeting() string {
	return fmt.Sprintf("Dear %s %s", u.Firstname, u.Lastname)
}

func main() {
	u := User{"Swastik", "Acharyya"}
	fmt.Println(u.Greeting())
}
