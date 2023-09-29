package wallet

import (
	"context"

	"github.com/DevShuxat/eater-service/src/domain/wallet/models"
	"github.com/DevShuxat/eater-service/src/domain/wallet/repositories"
	"gorm.io/gorm"
)

const (
	tableWallet = "eater.payment_cards"
)

type walletRepo struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) repositories.WalletRepository {
	return &walletRepo{
		db: db,
	}
}

func (r *walletRepo) AddCard(ctx context.Context, paymentCardR *models.PaymentCard)  error {
	result := r.db.WithContext(ctx).Table(tableWallet).Create(&paymentCardR)
	if result.Error != nil {
		return result.Error
	}
	return  nil
}

func (r *walletRepo) DeleteCard(ctx context.Context, cardID string) error {
	var paymentCard models.PaymentCard
	result := r.db.WithContext(ctx).Table(tableWallet).Delete(&paymentCard, "id = ?", cardID)
	if result.Error != nil {
		return nil
	}
	return nil
}

func (r *walletRepo) GetCard(ctx context.Context, paymentCardR string) (*models.PaymentCard,error){
	var paymentCard *models.PaymentCard
	result := r.db.WithContext(ctx).Table(tableWallet).First(&paymentCard, "id = ?", paymentCardR)
	if result.Error != nil {
		return nil, result.Error
	}
	return nil, nil
}

func (r *walletRepo) ListCardsByEater(ctx context.Context, eaterID string) ([]*models.PaymentCard, error) {
	var paymentCard []*models.PaymentCard
	result := r.db.
		WithContext(ctx).
		Table(tableWallet).
		Where("eater_id = ?", eaterID).
		Find(&paymentCard)

	if result.Error != nil {
		return nil, result.Error
	}
	return paymentCard, nil
}

