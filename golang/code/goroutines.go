package main

import (
	"fmt"
	"runtime"
	"sync"
)

// START OMIT
const Count = 10

func main() {
	var wg sync.WaitGroup
	wg.Add(Count)                // Tell the WaitGroup how many goroutines we're adding.
	for i := 0; i < Count; i++ { // Loop ten times to print concurrently.
		go func(n int) {
			fmt.Printf("goroutine#%d here! [total:%v]\n", n, runtime.NumGoroutine()) // HL
			wg.Done() // Tell the WaitGroup this goroutine is done.
		}(i + 1)
	}

	wg.Wait() // Wait for all the goroutines to finish.
}

// END OMIT
