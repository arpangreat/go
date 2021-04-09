package main

import (
	"fmt"
)

func main() {
	// ARRAYS

	// grades := [3]int{97, 85, 93}
	// fmt.Printf("Grades: %v", grades)

	// var students [3]string
	// fmt.Printf("Students: %v\n", students)
	// students[0] = "Sohini"
	// fmt.Printf("Students: %v\n", students)
	// fmt.Printf("Number of Students: %v", len(students))

	// var identityMatrix [3][3]int = [3][3]int{[3]int{1, 0, 0}, [3]int{0, 1, 0}, [3]int{0, 0, 1}}
	// fmt.Println(identityMatrix)

	// a := [...]int{1, 2, 3}
	// b := a
	// b[1] = 5
	// fmt.Println(a)
	// fmt.Println(b)
	// fmt.Println(len(a))

	// Slices
	a := []int{1, 2, 3}
	fmt.Println(a)
	fmt.Println(len(a))
}
