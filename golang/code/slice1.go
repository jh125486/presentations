package main

import (
	"fmt"
)

func main() {
    // START OMIT
    colors := []string{"Red", "Green", "Blue"}
    i := colors[1]
    fmt.Println(i)
    fmt.Println("Length:", len(colors))
    fmt.Println("Capacity:", cap(colors))
    // END OMIT
        
}
