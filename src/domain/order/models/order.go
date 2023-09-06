package models

import (
	"time"
)

type Order struct {
	ID            string
	EaterID       string
	Instruction   string
	RestaurantID  string // restaurant entity id
	Restaurant    *RestaurantInfo
	Delivery      *DeliveryInfo
	Payment       *PaymentInfo
	Items         []*OrderItem
	Status        string
	PaymentStatus string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
type OrderItem struct {
	ID         string
	ProductID  string
	Name       string
	Quantity   int
	Price      int
	TotalPrice int
	CreatedAt  time.Time
}
type RestaurantInfo struct {
	Name     string
	ImageUrl string
}
type AddressInfo struct {
	ID        string  `json:"id"`
	Name      string  `json:"name"`
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
}
type DeliveryInfo struct {
	Address *AddressInfo `json:"address"`
	Time    string       `json:"time"`
	Notes   string       `json:"notes"`
}
type PaymentInfo struct {
	Method        string
	CardID        string
	DeliveryPrice int
	TotalPrice    int
}
