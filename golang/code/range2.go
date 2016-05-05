package main

import "fmt"

func main() {
	// START OMIT
	fruits := map[string]int{
		"Apple":  45,
		"Mango":  24,
		"Orange": 34,
	}
    
    // Loop through keys and values
	for key, value := range fruits { // HL
		fmt.Println(key, value)
	} // HL

    // Loop through map keys
	for key := range fruits { // HL
		fmt.Println(key)
	} // HL

    // Loop through map values
	for _, value := range fruits { // HL
		fmt.Println(value)
	} // HL
	// END OMIT
}
