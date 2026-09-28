package main

import (
	"fmt"
	"sync"
)

// START OMIT
var (
	msg string
	wg  sync.WaitGroup
)

func setMessage() {
	defer wg.Done()
	msg = "Hello, World!"
}
func main() {
	wg.Add(1)
	go setMessage() // HL
	wg.Wait()
	fmt.Println(msg)
}

// END OMIT
