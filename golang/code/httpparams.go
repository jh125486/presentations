package main

import (
	"log"
	"net/http"
)

// helloHandler is an HTTP handler that writes a "hello, $BROWSER" response.
func helloHandler(w http.ResponseWriter, r *http.Request) {
	// Get the User-Agent from the HTTP request headers.
	userAgent := r.Header.Get("User-Agent")
	w.Write([]byte("Hello, client using " + userAgent))
}

func main() {
	// Set up the route for root.
	http.HandleFunc("/", helloHandler)

	// Start the server and listen on port 8080.
	log.Println("Starting server on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
