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
