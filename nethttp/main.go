package main

import (
	"fmt"
	"net/http"
)

// handler is a typical HTTP request-response handler in Go; details later
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
