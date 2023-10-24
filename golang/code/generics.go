package main

import (
	"fmt"
)

// START OMIT
type number interface {
	float64 | int | int64 // constraints
}

func Sum[T number](slice []T) T { // HL
	var total T
	for _, v := range slice {
		total += v
	}
	return total
}

func main() {
	fmt.Println("Ints:", Sum([]int{512, 50, 23, 30, 7}))
	fmt.Println("Floats:", Sum([]float64{0.061, 0.097, 0.085}))
}

// END OMIT
