package models

import (
	"time"
)

type DeliveryRating struct {
	ID        string    `json:"id" gorm:"primaryKey"`
	EaterID   string    `json:"eater_id" gorm:"index:idx_d_rating_by_eater"`
	OrderID   string    `json:"order_id" gorm:"index:idx_d_rating_by_order"`
	Rating    int       `json:"rating"` // 1-5
	Comment   string    `json:"comment"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
