package model

type Order struct {
	Id     int   `json:"id"`
	Price  int32 `json:"price"`
	TourID int16 `json:"tour_id"`
}
