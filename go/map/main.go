package main

import (
	"fmt"
)

type Point struct {
	X int
	Y int
}

func main() {
	m := map[int]int32{
		1: 0,
	}

	e, ok := m[2]
	fmt.Println(e)
	fmt.Println(ok)
}
