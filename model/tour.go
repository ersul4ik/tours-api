package model

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Tour struct {
	Id          int       `json:"id"`
	Title       string    `json:"title"`
	Content     string    `json:"content"`
	DepartureAt time.Time `json:"departure_at"`
}

func (u Tour) String() string {
	return fmt.Sprintf("Tour<%d %s>", u.Id, u.Title)
}

func GetTours(db *gorm.DB) ([]Tour, error) {
	tours := []Tour{}
	result := db.Find(&tours)
	return tours, result.Error
}

func GetTour(db *gorm.DB, id int) (Tour, error) {
	tour := Tour{}
	result := db.First(&tour, "id = ?", id)
	return tour, result.Error
}

func CreateTour(db *gorm.DB, tour Tour) error {
	fmt.Printf("ID: %d", tour.Id)
	result := db.Create(&tour)
	return result.Error
}
