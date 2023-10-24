package main

import "fmt"

func main() {
	// START OMIT
	defer fmt.Println("again")
	defer fmt.Println("world")
	fmt.Println("hello")
	// END OMIT
}
