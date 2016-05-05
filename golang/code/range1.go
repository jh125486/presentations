package main

import "fmt"

func main() {
	// START OMIT
	var fruits = []string{"Apple", "Mango", "Orange"}
	
    for index, value := range fruits { // HL
		fmt.Println(index, value)
	} // HL
	// END OMIT
}
