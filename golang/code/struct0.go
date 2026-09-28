package main

import "fmt"

// START OMIT
type Point struct {
	X, Y int
}

func main() {
	var origin Point // zero value: {0 0}
	p := Point{X: 3, Y: 4}
	fmt.Println(origin, p, p.X, p.Y)
}

// END OMIT
