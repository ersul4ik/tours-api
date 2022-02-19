package main

import (
	"fmt"
	"log"
	"net/http"
)

func getTours(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
}

func handleRequests() {
	http.HandleFunc("/", getTours)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func main() {
	handleRequests()
}
