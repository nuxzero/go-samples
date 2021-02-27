package main

import "fmt"

type Vertex5 struct {
	Lat, Long float64
}

var m map[string]Vertex5

func main() {
	m = make(map[string]Vertex5)
	m["Bell labs"] = Vertex5{40.68433, -74.39967}
	fmt.Println(m["Bell labs"])
}
