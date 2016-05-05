package main

import "fmt"

func main() {
	// START OMIT
	i := 0 // initialized outside loop
    
	for {  // To infinity! // HL
		if i >= 5 {
			break
		}
		i++
		fmt.Println("Number:", i)
	} // HL
	// END OMIT
}
