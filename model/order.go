package model

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Order struct {
	Id      int `json:"id"`
	TourID  int `json:"tour_id" validate:"required"`
	Tour    Tour
	Created time.Time `gorm:"autoCreateTime" json:"created"`
}

func (o Order) String() string {
	return fmt.Sprintf("Order<%d>", o.Id)
}

func CreateOrder(db *gorm.DB, order Order) error {
	fmt.Printf("ID: %d", order.Id)
	result := db.Create(&order)
	return result.Error
}

func GetOrders(db *gorm.DB) ([]Order, error) {
	orders := []Order{}
	result := db.Find(&orders)
	return orders, result.Error
}
