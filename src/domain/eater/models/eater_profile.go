package models

import (
	"time"
)

type EaterProfile struct {
	EaterID     string
	Name        string
	PhoneNumber string
	ImageUrl    string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
