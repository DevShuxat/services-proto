package repositories

import (
	"context"

	"github.com/DevShuxat/eater-service/src/domain/wallet/models"
)

type WalletRepository interface {
	WithTx(ctx context.Context, f func(r WalletRepository) error) error
	AddCard(ctx context.Context, CardToken string, Number string) (*models.PaymentCard, error)
	GetCard(ctx context.Context, ID string) error
	DeleteCard(ctx context.Context, ID string) error
	ListCard(ctx context.Context, ID string) (*models.PaymentCard, error)
	ListCardsByEater(ctx context.Context, eaterID string) ([]*models.PaymentCard, error)
}
