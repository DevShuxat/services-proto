package models

import (
	"time"
)

type Eater struct {
	ID           string
	PhoneNumber  string
	PasswordHash string
	PasswordSalt string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
