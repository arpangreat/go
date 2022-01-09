package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello")

	var whatToSay string = "Bye, Cruel World"
	var i int = 12

	fmt.Println(whatToSay)

	fmt.Println("i is set to", i)

	whatWasSaid, theOtherThing := saySomething()

	fmt.Println("The functioned returned ", whatWasSaid, theOtherThing)
}

func saySomething() (string, string) {
	return "something", "else"
}
