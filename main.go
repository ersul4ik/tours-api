package main

import (
	"encoding/json"
	"fmt"
	"log"
	"main/config"
	"main/database"
	"main/model"
	"net/http"
	"strconv"
	"time"

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

// getTour returns tour by given ID
func getTour(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Fatalf("Converting process is failed: %s", err)
	}

	fmt.Printf("Look at the tour with ID: %d", id)
	tour, err := model.GetTour(db, id)
	if err != nil && err.Error() != "pg: no rows in result set" {
		log.Fatalf("GetTour: %s", err)
	}
	json.NewEncoder(w).Encode(tour)
}

// handleRequests manages requests
func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/tours", getTours)
	router.HandleFunc("/tours/{id:[0-9]+}/", getTour)

	srv := &http.Server{
		Addr:         "0.0.0.0:10000",
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      router,
	}
	fmt.Println("\nServer is running ... ")
	log.Fatal(srv.ListenAndServe())
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
