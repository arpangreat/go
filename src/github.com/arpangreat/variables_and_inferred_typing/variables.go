package main

import (
	"fmt"
)

var (
	name, location string = "Swastik Acharyya", "Rampurhat"
	age            int    = 20
)

func main() {
	fmt.Printf("%s (%d) of %s", name, age, location)
}
