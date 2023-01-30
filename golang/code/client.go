package main

import (
    "bufio"
    "fmt"
    "net/http"
)

// START OMIT

func main() {
    resp, _ := http.Get("https://digital-divas.weebly.com/") // omit error handling
    defer resp.Body.Close()    
	fmt.Println("Response status:", resp.Status)

	scanner := bufio.NewScanner(resp.Body)
    for i := 0; scanner.Scan() && i < 5; i++ {
        fmt.Println(scanner.Text())
    }
}

// END OMIT