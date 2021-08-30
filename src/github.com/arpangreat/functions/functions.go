package main

import (
	"fmt"
)

func location(city string) (string, string) {
	var region, continent string

	switch city {
	case "Rampurhat", "Bolpur":
		region, continent = "Birbhum", "Asia"
	case "Kolkata", "Alipur":
		region, continent = "Kolkata", "Asia"
	default:
		region, continent = "Unknown", "Unknown"
	}
	return region, continent
}

func main() {
	region, continent := location("Rampurhat")
	fmt.Printf("Swastik is living in: %s (%s)\n", region, continent)
}
