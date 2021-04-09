package main

import (
	"fmt"
)

func main() {
	// name, location := "Swastik", "Rampurhat"
	// age := 20
	// fmt.Printf("%s (%d) of %s", name, age, location)

	fmt.Println(add(42, 23))
	region, continent := location("Rph")
	fmt.Printf("Swastik lives in %s, %s", region, continent)
}

func add(x int, y int) int {
	return x + y
}

func location(city string) (string, string) {
	var region string
	var continent string
	switch city {
	case "Rampurhat", "Rph", "Tarapith":
		region, continent = "West Bengal", "Asia"
	case "Kurumgram", "Kmg":
		region, continent = "Birbhum", "Asia"
	default:
		region, continent = "West Bengal", "Asia"
	}
	return region, continent
}
