package main

import "fmt"

func main() {
	var a [2]string
	a[0] = "Swastik"
	a[1] = "Acharyya"

	fmt.Println(a[0], a[1])
	fmt.Println(a)

	b := [2]string{"string", "used"}
	fmt.Printf("%q\n", b)

	// Multi-dimensional Arrays
	var m [2][3]string
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			m[i][j] = fmt.Sprintf("row %d - column %d\n", i+1, j+1)
		}
	}

	fmt.Printf("%q\n", m)
	// [["row 1 - column 1" "row 1 - column 2" "row 1 - column 3"]
	//  ["row 2 - column 1" "row 2 - column 2" "row 2 - column 3"]]
}
