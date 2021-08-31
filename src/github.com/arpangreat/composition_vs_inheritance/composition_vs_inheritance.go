package main

import (
	"fmt"
)

type User struct {
	Id       int
	Name     string
	Location string
}

func (u *Player) Greetings() string {
	return fmt.Sprintf("Hi %s from %s", u.Name, u.Location)
}

type Player struct {
	Id       int
	Name     string
	Location string
	GameId   int
}

func main() {
	p := Player{}
	p.Name = "Swastik"
	p.Id = 20
	p.Location = "Rampurhat, Birbhum"
	p.GameId = 90390
	fmt.Printf("%+v\n", p)
	fmt.Println(p.Greetings())
}
