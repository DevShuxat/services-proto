package models

import (
	"time"
)

type EaterSmsCode struct {
	ID        int       `json:"id" gorm:"primaryKey"`
	EaterID   string    `json:"eater_id" gorm:"index:idx_sms_code_by_eater"`
	Code      string    `json:"code"`
	ExpiresIn int       `json:"expires_in"`
	CreatedAt time.Time `json:"created_at"`
}
