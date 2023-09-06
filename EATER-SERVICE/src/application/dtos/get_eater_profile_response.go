package dtos

import (
	"time"
	
)

type GetEaterProfileResponse struct {
	Name string `json:"name" validate:"required"`
	PhoneNumber string `json:"phone_number" validate:"required"`
	ImageUrl string `json:"image_url" validate:"required"`
	CreatedAt time.Time `query:"start" validate:"required"`
	UpdatedAt time.Time 
}
	