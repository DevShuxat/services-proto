package models

import (
	"time"
)

type RestaurantRating struct {
	ID           string `json:"id" gorm:"primaryKey"`
	EaterID      string `json:"eater_id" gorm:"index:idx_r_rating_by_eater"`
	RestaurantID string `json:"restaurant_id" gorm:"index:idx_r_rating_by_restaurant"` //entity       id
	Rating       int       `json:"rating"` //1-5
	Comment      string    `json:"comment"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
