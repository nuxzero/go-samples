package main

import "fmt"

var pow1 = []int{1, 2, 4, 8, 12, 32, 64, 128}

func main() {
	for i, v := range pow1 {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}
