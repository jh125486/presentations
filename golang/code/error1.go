package main

import (
	"fmt"
	"strconv"
)

func main() {
	// START OMIT
	i, err := strconv.Atoi("42")
	if err != nil { // HL
		fmt.Printf("couldn't convert number: %v\n", err)
		return
	}
	fmt.Println("Converted integer:", i)

	// Should fail to convert
	if i, err := strconv.Atoi("Elephant"); err != nil { // HL
		fmt.Printf("couldn't convert number: %v\n", err)
		return
	} else {
		fmt.Println("Converted integer:", i)
	}
	// END OMIT
}
