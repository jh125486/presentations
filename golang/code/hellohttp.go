package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// helloHandler is an HTTP handler that writes a "hello, world" response.
func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World! The time is %s.", time.Now().Format(time.Kitchen))
}

func main() {
	// Set up the route for root.
	http.HandleFunc("/", helloHandler)

	// Start the server and listen on port 8080.
	log.Println("Starting server on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
