package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	// Get user's home directory.
	dirname, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	// Serve static files from the home directory using the file server handler.
	http.Handle("/", http.FileServer(http.Dir(dirname)))

	// Start the server on port 8080.
	log.Println("Starting server on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
