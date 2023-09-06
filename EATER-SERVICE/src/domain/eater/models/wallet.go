package models

import (
	"time"
)

type PaymentCard struct {
	ID         string
	EaterID    string
	Number     string
	CardToken  string // we don't need a card token here but only in Payment Microservice in
	Restaurant string
	IsVerifed  bool
	CreatedAt  time.Time
}
