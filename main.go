package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/", handler) // each request calls handler
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// handler echoes the Path component of the requested URL.
func handler(w http.ResponseWriter, r *http.Request) {
	path := strings.Split(r.URL.Path, "/")[1:]
	fmt.Println(path)
	if len(path) > 1 && path[0] == "greet" {

	}
	fmt.Fprintf(w, "Hello, %v\n", path[1])
}
