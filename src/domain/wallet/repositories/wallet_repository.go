package repositories

import (
	"context"

	"github.com/DevShuxat/eater-service/src/domain/wallet/models"
)

type WalletRepository interface {
	AddCard(ctx context.Context, paymentCardR *models.PaymentCard)  error
	GetCard(ctx context.Context, paymentCardR string) ([]*models.PaymentCard,error)
	DeleteCard(ctx context.Context, cardID string) error
	ListCard(ctx context.Context, cardID string, sort string, page, pageSize int) ([]*models.PaymentCard, error)
	ListCardsByEater(ctx context.Context, eaterID string) ([]*models.PaymentCard, error)
}
