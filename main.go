package main

import (
	"encoding/json"
	"log"
	"main/model"
	"net/http"
)

var Tours []model.Tour

func getTours(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(Tours)
}

func handleRequests() {
	http.HandleFunc("/", getTours)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func main() {
	Tours = []model.Tour{
		{Id: 1, Title: "Foo #1", Content: "Content #1"},
		{Id: 2, Title: "Bar #1", Content: "Content #2"},
	}
	handleRequests()
}
