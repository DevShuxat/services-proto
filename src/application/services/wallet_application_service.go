package services
import (
	"context"

	"github.com/DevShuxat/eater-service/src/domain/wallet/models"
)

type WalletRepository interface {
	AddCard(ctx context.Context, cardNumber string, CardToken string) (*models.PaymentCard, error)
	GetCard(ctx context.Context, cardID string) error
	DeleteCard(ctx context.Context, cardID string) error
	ListCard(ctx context.Context, cardID string) (*models.PaymentCard, error)
}
