package dtos

import (
	"time"
)

type ConfirmSMSCodeResponse struct {
	Code      string
	EaterID   string
	CreatedAt time.Time
	UpdatedAt time.Time
}
