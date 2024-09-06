package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func setupHttpHandlers() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Write to the console.
		fmt.Printf("%s\n", html.EscapeString(r.URL.Path))
		// Write the response.
		fmt.Fprintf(w, "%q\n", html.EscapeString(r.URL.Path))
	})
}
func main() {
	setupHttpHandlers()
	log.Fatal(http.ListenAndServe(":8080", nil))
}
