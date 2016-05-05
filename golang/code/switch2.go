package main

import (
	"fmt"
	"time"
)

func main() {
	// START OMIT
	t := time.Now()
    
	switch { // HL
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	
    case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	
    default:
		fmt.Println("Good evening.")
	} // HL
	// END OMIT
}
