package main

import (
	"fmt"
)

var (
	i int     = 42
	f float64 = float64(i)
	u uint64  = uint64(f)
)

func main() {
	fmt.Printf("%d , %f, %d", i, f, u)
}
