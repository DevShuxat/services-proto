package wallet

import (
	"context"

	"github.com/DevShuxat/eater-service/src/domain/wallet/models"
	"github.com/DevShuxat/eater-service/src/domain/wallet/repositories"
	"github.com/DevShuxat/eater-service/src/infrastructure/utils"
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

func (r *walletRepo) AddCard(ctx context.Context, CardToken string, Number string) ([]*models.PaymentCard, error) {
	result := r.db.WithContext(ctx).Table(tableWallet).Create(&CardToken)
	if result.Error != nil {
		return nil, result.Error
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

func (r *walletRepo) GetCard(ctx context.Context, cardID string) error {
	var paymentCard *models.PaymentCard
	result := r.db.WithContext(ctx).Table(tableWallet).First(&paymentCard, "id = ?", cardID)
	if result.Error != nil {
		return result.Error
	}
	return nil
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
func (r *walletRepo) ListCard(ctx context.Context, cardID string, sort string, page, pageSize int) ([]*models.PaymentCard, error) {
	var paymentCard []*models.PaymentCard
	result := r.db.
		WithContext(ctx).
		Table(tableWallet).
		Where("eater_id = ?", cardID).
		Scopes(utils.Paginate(page, pageSize), utils.Sort(sort)).Find(&paymentCard)

	if result.Error != nil {
		return nil, result.Error
	}
	return paymentCard, nil
}
