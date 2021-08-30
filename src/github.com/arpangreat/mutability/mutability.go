package main

import (
	"fmt"
)

type Artist struct {
	Name, Genre string
	Songs       int
}

func newRelease(a *Artist) int {
	a.Songs++
	return a.Songs
}

func main() {
	me := &Artist{Name: "Swastik", Genre: "Electro", Songs: 11}
	fmt.Printf("%s released their %dth song\n", me.Name, newRelease(me))
	fmt.Printf("%s has total %d song\n", me.Name, me.Songs)
}
