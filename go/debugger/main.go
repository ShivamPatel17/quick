package main

import "fmt"

func main() {
	// also need to test this out with uuid and other packages
	a := 2
	b := 2

	for i := range 5 {
		a = b + i
		b = a + 1
	}
	fmt.Println(a)
}
