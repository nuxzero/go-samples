package main

import "fmt"

func main() {
	// Declare array with zero value
	var arr [3]int
	fmt.Printf("arr: %v\n", arr) // arr: [0 0 0]

	// Declare and initialize array with literal syntax
	arr2 := [3]string{"a", "b", "c"}
	fmt.Printf("arr2: %v\n", arr2) // arr2: [a b c]

	// Declare slice with zero value
	var sl []int
	fmt.Printf("sl: %v\n", sl) // sl: []

	// Declare and initialize slice with literal syntax
	sl2 := []int{1, 2, 3}
	fmt.Printf("sl2: %v\n", sl2)               // sl2: [1 2 3]
	fmt.Printf("ls2 length: %v\n", len(sl2))   // ls2 length: 3
	fmt.Printf("ls2 capacity: %v\n", cap(sl2)) // ls2 capacity: 3

	// Two dimensional slice
	sl3 := [][]int{{1, 2}, {3, 4}}
	fmt.Printf("sl3: %v\n", sl3) // sl3: [[1 2] [3 4]]
}
