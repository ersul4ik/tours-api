package model

import (
	"fmt"

	"github.com/go-pg/pg/v10"
)

type Tour struct {
	Id      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

func (u Tour) String() string {
	return fmt.Sprintf("Tour<%d %s>", u.Id, u.Title)
}

func GetTours(db *pg.DB) ([]Tour, error) {
	var tours []Tour
	err := db.Model(&tours).Order("id ASC").Limit(20).Select()
	return tours, err
}

func GetTour(db *pg.DB, id int) (Tour, error) {
	tour := Tour{}
	err := db.Model(&tour).Where("id = ?", id).Select()
	return tour, err
}

func CreateOrUpdate(db *pg.DB, tour Tour) error {
	fmt.Printf("ID: %d", tour.Id)
	_, err := db.Model(&tour).
		OnConflict("(id) DO UPDATE").
		Insert()
	return err
}
