package main

import (
	"fmt"
)

func main() {
	if err := foo(); err != nil {
		panic(err)
	}
}
