package main

import "fmt"

func main() {
	// START OMIT
	i := 0      // initialized outside loop
    
	for i < 5 { // HL
		i++
		fmt.Println("Number:", i)
	} // HL
	// END OMIT
}
