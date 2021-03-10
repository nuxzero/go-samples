package main

import (
	"fmt"
	"math"
)

type Vertex9 struct {
	X, Y float64
}

func Abs(v Vertex9) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex9{3, 4}
	fmt.Println(Abs(v))
}
