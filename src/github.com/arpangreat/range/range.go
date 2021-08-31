package main

import (
	"fmt"
)

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	// With Make()
	fow := make([]int, 10)
	for i := range fow {
		fow[i] = 1 << uint(i)
	}
	for _, value := range fow {
		fmt.Printf("%d\n", value)
	}

	// Break & Continue
	for i := range fow {
		fow[i] = 1 << uint(1)
		if fow[1] >= 16 {
			break
		}
	}
	fmt.Println(fow)

	// Range and Maps
	cities := map[string]int{
		"Rampurhat": 731224,
		"Bolpur":    677812,
		"Suri":      124589,
	}
	for key, value := range cities {
		fmt.Printf("%s has %d inhabitants\n", key, value)
	}
}
