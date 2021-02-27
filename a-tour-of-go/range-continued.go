package main

import "fmt"

func main() {
	pow2 := make([]int, 10)
	for i := range pow2 {
		pow2[i] = 1 << uint(i)
	}
	for _, value := range pow2 {
		fmt.Printf("%d\n", value)
	}
}
