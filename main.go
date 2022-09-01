package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"main/config"
	"main/model"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

// getTours returns tours existing in DB
func getTours(w http.ResponseWriter, r *http.Request) {
	tours, err := model.GetTours(db)
	if err != nil {
		log.Fatalf("GetTours: %s", err)
	}

	json.NewEncoder(w).Encode(tours)
}

// createTour allow to create a tour instance
func createTour(w http.ResponseWriter, r *http.Request) {
	var tour model.Tour

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("readAll error")
		panic(err)
	}

	err = json.Unmarshal(body, &tour)
	if err != nil {
		fmt.Println("Unmarshal error")
		panic(err)
	}
	json.NewEncoder(w).Encode(tour)
	err = model.CreateTour(db, tour)
	if err != nil {
		log.Fatalf("Error during creating the tour. err: %s", err)
	}
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
	if err != nil {
		log.Fatalf("GetTour: %s", err)
	}

	json.NewEncoder(w).Encode(tour)
}

// getTour returns tour by given ID
func getOrders(w http.ResponseWriter, r *http.Request) {
	log.Print("GetOrders: request is receiving")

	orders, err := model.GetOrders(db)
	if err != nil {
		log.Fatalf("GetOrders: %s", err)
	}

	json.NewEncoder(w).Encode(orders)
}

func createOrder(w http.ResponseWriter, r *http.Request) {
	var order model.Order

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("readAll error")
		panic(err)
	}

	err = json.Unmarshal(body, &order)
	if err != nil {
		fmt.Println("Unmarshal error")
		panic(err)
	}
	err = model.CreateOrder(db, order)
	if err != nil {
		log.Fatalf("Error during creating the order. err: %s", err)
		return
	}
	json.NewEncoder(w).Encode(order)
}

// handleRequests manages requests
func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/tours/", getTours).Methods("GET")
	router.HandleFunc("/tours/", createTour).Methods("POST")
	router.HandleFunc("/tours/{id:[0-9]+}/", getTour)
	router.HandleFunc("/orders/", getOrders).Methods("GET")
	router.HandleFunc("/orders/", createOrder).Methods("POST")

	host := "0.0.0.0:10000"
	srv := &http.Server{
		Addr:         host,
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      router,
	}
	log.Printf("\nServer is running:  %s", host)
	log.Fatal(srv.ListenAndServe())
}

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Fatal("No .env file found")
	}
}

func main() {
	conf := config.New()
	database, err := gorm.Open(postgres.Open(conf.DATABASE_URL), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	database.AutoMigrate(&model.Tour{})
	database.AutoMigrate(&model.Order{})
	db = database
	handleRequests()
}
