package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
)

func main() {
	num := rand.Intn(10) + 1
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		guess, _ := strconv.Atoi(r.URL.Query().Get("guess")) 
		switch {
		case guess < num:
			fmt.Fprint(w, "Too low!")
		case guess > num:
			fmt.Fprint(w, "Too high!")
		default:
			fmt.Fprint(w, "You got it!")
		}
	})
	http.ListenAndServe(":8080", nil)
}
