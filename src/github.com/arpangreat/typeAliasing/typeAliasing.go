package main

import (
	"fmt"
	"strings"
)

type MyStr string

func (s MyStr) Uppercase() string {
	return strings.ToUpper(string(s))
}

func main() {
	fmt.Println(MyStr("test").Uppercase())
}
