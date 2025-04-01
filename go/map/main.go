package main

import (
	"fmt"
)

type Point struct {
	X int
	Y int
}

func main() {
	m := make(map[Point]Point)

	p1 := Point{X: 3, Y: 4}

	p2 := m[p1]
	p2.X = 10
	p2.Y = 1

	fmt.Println(p2) // Output: Point 1
}
