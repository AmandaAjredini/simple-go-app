package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	log.Println("Received a request at /")
	fmt.Fprintf(w, "Hello from OpenShift!")
}

func main() {
	http.HandleFunc("/", handler)
	log.Println("Starting server on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
