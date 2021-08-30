package main

import (
	"fmt"
	"time"
)

type Stringer interface {
	String() string
}

type fakeString struct {
	content string
}

// function used to implement the Stringer interface
func (s *fakeString) String() string {
	return s.content
}

func printString(value interface{}) {
	switch str := value.(type) {
	case string:
		fmt.Println(str)
	case Stringer:
		fmt.Println(str.String())
	}
}

func timeMap(y interface{}) {
	z, ok := y.(map[string]interface{})
	if ok {
		z["updated_at"] = time.Now()
	}
}

func main() {
	foo := map[string]interface{}{
		"Swastik": 42,
	}
	timeMap(foo)
	fmt.Println(foo)

	s := &fakeString{"Ceci n'est pas un string"}
	printString(s)
	printString("Hello, Gophers")
}
