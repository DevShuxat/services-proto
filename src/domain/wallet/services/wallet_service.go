package services

import (
	"context"
	"time"

	"github.com/DevShuxat/eater-service/src/domain/wallet/models"
	"github.com/DevShuxat/eater-service/src/domain/wallet/repositories"
	"github.com/DevShuxat/eater-service/src/infrastructure/rand"
	"github.com/DevShuxat/eater-service/src/infrastructure/utils"
)

type WalletService interface {
	AddCard(ctx context.Context, CardToken, EaterID string, Number string,IsVerified bool) ([]*models.PaymentCard, error)
	GetCard(ctx context.Context, cardID string) (*models.PaymentCard, error)
	DeleteCard(ctx context.Context, cardID string) error
	ListCard(ctx context.Context, cardID string, sort string, page, pageSize int) ([]*models.PaymentCard, error)
	ListCardsByEater(ctx context.Context, eaterID string) ([]*models.PaymentCard, error)
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

func (s *paymentSvcImpl) 	AddCard(ctx context.Context, CardToken, EaterID string, Number string,IsVerified bool) ([]*models.PaymentCard, error){
	paymentCardR :=  models.PaymentCard {
		ID: rand.UUID(),
		EaterID: EaterID,
		Number: Number,
		IsVerified: true,
		CreatedAt: time.Now().UTC(),

	}
	 err := s.walletRepo.AddCard(ctx, &paymentCardR)
	if err != nil {
		return nil, err
	}
	return []*models.PaymentCard{&paymentCardR}, nil
}

func (s *paymentSvcImpl) DeleteCard(ctx context.Context, cardID string) error {
	err := s.walletRepo.DeleteCard(ctx, cardID)
	if err != nil {
		return err
	}
	return nil
}

func (s *paymentSvcImpl) GetCard(ctx context.Context, cardID string) (*models.PaymentCard, error) {
	var paymentCardR models.PaymentCard
	err := s.walletRepo.GetCard(ctx, &paymentCardR)
	if err != nil {
		return nil, err
	}
	return &paymentCardR, nil
}

func (s *paymentSvcImpl) ListCardsByEater(ctx context.Context, eaterID string) ([]*models.PaymentCard, error) {
	paymentCards, err := s.walletRepo.ListCardsByEater(ctx, eaterID)
	if err != nil {
		return nil, err
	}
	return paymentCards, nil
}

func (s *paymentSvcImpl) ListCard(ctx context.Context, cardID string, sort string, page, pageSize int) ([]*models.PaymentCard, error) {
	result := s.walletRepo.ListCard(ctx, cardID)
	if err != nil {
		return nil, err
	}
	result = result.Scopes(utils.Sort(sort))
	result = result.Scopes(utils.Paginate(page, pageSize))

	paymentCards, err := result.FindPaymentCards()
	if err != nil {
		return nil, err
	}
	return paymentCards, nil
}