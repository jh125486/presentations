package main

import (
	"fmt"
)

// START OMIT
type Sounder interface {
	Sound() string
}

type Dog struct{}

func (d Dog) Sound() string {
	return "Woof"
}

// Receive an interface, return a concrete type.
func AnimalSound(s Sounder) string {
	return s.Sound()
}

func main() {
	fmt.Println(AnimalSound(Dog{}))
}

// END OMIT
