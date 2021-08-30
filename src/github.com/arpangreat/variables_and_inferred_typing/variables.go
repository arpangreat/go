package main

import (
	"fmt"
)

var (
	name, location string = "Swastik Acharyya", "Rampurhat"
	age            int    = 20
)

const (
	StatusOk        = 200
	StatusCreated   = 201
	StatusAccepted  = 202
	StatusNoContent = 204
)

func main() {
	action := func() {
		fmt.Printf("%d , %d, %d, %d", StatusOk, StatusCreated, StatusAccepted, StatusNoContent)
	}

	fmt.Printf("%s (%d) of %s", name, age, location)
	action()
}
