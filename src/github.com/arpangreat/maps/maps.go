package main

import (
	"fmt"
)

// Make Type of Map with Go Conventions
type Vertex struct {
	Lat, Long float64
}

var a map[string]Vertex

// or

var m = map[string]Vertex{
	"Bell Labs": {40.68433, -74.39967},
	// same as "Bell Labs": Vertex{40.68433, -74.39967}
	"Google": {37.42202, -122.08408},
}

func main() {
	celebs := map[string]int{
		"Swastik":      20,
		"Tj Devries":   30,
		"ThePrimeagen": 40,
	}
	fmt.Printf("%#v\n", celebs)

	// Maps with Go Conventions
	m = make(map[string]Vertex)
	m["Bar labs"] = Vertex{40.3411, -38.9012}
	fmt.Println(m["Bar labs"])
}
