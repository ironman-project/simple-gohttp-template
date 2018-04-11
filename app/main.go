package main

import (
	"fmt"
	"log"
	"net/http"
)

func registerEndpoint(path string, h http.HandlerFunc) error {
	http.HandleFunc(path, h)
	return nil
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %s", "{{.Values.myName | default "World"}}")
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
