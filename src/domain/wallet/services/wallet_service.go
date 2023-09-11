package services

import (
	"context"
	"fmt"
	"time"

	"github.com/DevShuxat/eater-service/src/domain/wallet/models"
	"github.com/DevShuxat/eater-service/src/domain/wallet/repositories"
	"github.com/DevShuxat/eater-service/src/infrastructure/rand"
)

type WalletRepository interface {
	WithTx(ctx context.Context, f func(r WalletRepository) error) error
	AddCard(ctx context.Context, CardToken string, Number string) (*models.PaymentCard, error)
	GetCard(ctx context.Context, ID string) error
	DeleteCard(ctx context.Context, ID string) error
	ListCard(ctx context.Context, ID string) (*models.PaymentCard, error)
}

type paymentSvcImpl struct {
	walletRepo repositories.WalletRepository
}

func NewWalletService(
	walletRepo repositories.WalletRepository,
) WalletService {
	return &paymentSvcImpl{
		walletRepo: walletRepo,
	}
}
func (s *paymentSvcImpl) AddCard(ctx context.Context, CardToken string, Number string) (*models.PaymentCard, error) {
	isVerified := true

	card, err := s.walletRepo.GetCard(ctx, Number)
	if err != nil {
		isVerified = false
	}

	if isVerified {
		return s.handleExistingWallet(ctx, CardToken)
	}
	return s.handleNewWallet(ctx, Number, CardToken)

}

func (s *paymentSvcImpl) handleNewWallet(ctx context.Context, Number string, CardToken string) (string, error) {
	var (
		id         = rand.UUID()
		eaterID    = rand.UUID()
		number     = fmt.Sprintf("eater-%s", rand.NumericString(5))
		cardToken  = fmt.Sprintf("card-%s", rand.NumericString(10))
		isVerified = true
		now        = time.Now().UTC()
	)
	payment := &models.PaymentCard{
		ID:         id,
		EaterID:    eaterID,
		Number:     number,
		CardToken:  cardToken,
		IsVerified: isVerified,
		CreatedAt:  now,
	}

	err := s.walletRepo.WithTx(ctx, func(r repositories.WalletRepository) error {
		if _, err := r.AddCard(ctx, Number, CardToken); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return "", err
	}
	return eaterID, nil

}

func (s *paymentSvcImpl) handleExistingWallet(ctx context.Context, CardToken string) (string, error) {
	eater, err := s.walletRepo.GetCard(ctx, CardToken)
	if err != nil {
		return "", err
	}
	return cardToken, nil

}

func (s *paymentSvcImpl) GetCard(ctx context.Context, ID string) (*models.PaymentCard, error) {
	card, err := s.walletRepo.GetCard(ctx, ID)
	if err != nil {
		return nil, err
	}
	return card, nil
}
