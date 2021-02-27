package main

import "fmt"

func fibonacci() func() int {
	return func() int {

	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
