package main

import (
	"encoding/json"
	"log"
	"main/config"
	"main/database"
	"main/model"
	"net/http"

	"github.com/go-pg/pg/v10"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

var db *pg.DB

// getTours returns tours existing in DB
func getTours(w http.ResponseWriter, r *http.Request) {
	tours, err := model.GetTours(db)
	if err != nil {
		log.Fatalf("GetTours: %s", err)
	}

	json.NewEncoder(w).Encode(tours)
}

// handleRequests manages requests
func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/tours", getTours)

	log.Fatal(http.ListenAndServe(":10000", router))
}

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	conf := config.New()
	opt, err := pg.ParseURL(conf.DATABASE_URL)
	if err != nil {
		panic(err)
	}

	db = pg.Connect(opt)
	defer db.Close()

	err = database.CreateSchema(db)
	if err != nil {
		panic(err)
	}

	handleRequests()
}
