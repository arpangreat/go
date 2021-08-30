package main

import (
	"fmt"
	"time"
)

type Bootcamp struct {
	lat, lon float64
	Date     time.Time
}

type Point struct {
	X, Y int
}

var (
	p = Point{1, 2}
	q = &Point{1, 2}
	r = Point{X: 1}
	s = Point{}
)

func main() {
	fmt.Println(Bootcamp{
		lat:  34.00021,
		lon:  -118.90267,
		Date: time.Now(),
	})

	fmt.Println(p, q, r, s)

	event := Bootcamp{
		lat: 32.90021,
		lon: -124.80345,
	}
	event.Date = time.Now()

	fmt.Printf("Event on %s, location (%f, %f)",
		event.Date, event.lat, event.lon)
}
