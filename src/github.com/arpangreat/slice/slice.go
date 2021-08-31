package main

import (
	"fmt"
)

func main() {
	p := []int{2, 3, 4, 5, 67, 81}
	fmt.Println(p)
	fmt.Println(p[1:4])
	fmt.Println(p[:3])
	fmt.Println(p[4:])

	// making a slice using make() func
	cities := make([]string, 3)
	cities[0] = "Rampurhat"
	cities[1] = "Bolpur"
	cities[2] = "Suri"
	fmt.Printf("%q\n", cities)

	// Appending to slice
	schools := []string{}
	schools = append(schools, "RJLV")
	schoolsAndcities := append(schools, cities...)
	fmt.Println(schools)
	fmt.Println(schoolsAndcities)

	// Length
	countries := make([]string, 42)
	fmt.Println(len(countries))

	// NIL slices
	var z []int
	fmt.Println(z, len(z), cap(z))

	if z == nil {
		fmt.Println("nill!")
	}
}
