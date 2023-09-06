package dtos

import (
	"time"
)

type EaterSignupResponse struct {
	PhoneNumber string `json:"phone_number"`
	PasswordHash string `json:"password_hash"`
	PasswordSalt string `json:"password_salt"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
