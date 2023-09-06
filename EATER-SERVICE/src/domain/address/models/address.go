package models

import (
	"time"
)

type Address struct {
	ID string `json:"id" gorm:"primaryKey"`
	EaterID string `json:"eater_id"`
	Name string `json:"name"`
	Location *Location `json:"location"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
