package main

import (
	"fmt"
)

type User struct {
	Firstname, Lastname string
}

func (u *User) Name() string {
	return fmt.Sprintf("%s %s", u.Firstname, u.Lastname)
}

type Namer interface {
	Name() string
}

func Greet(n Namer) string {
	return fmt.Sprintf("Dear %s", n.Name())
}

func main() {
	u := &User{"Swastik", "Acharyya"}
	fmt.Println(Greet(u))
}
