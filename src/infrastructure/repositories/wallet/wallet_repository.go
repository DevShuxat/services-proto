package wallet

import (
	"context"

	"github.com/DevShuxat/eater-service/src/domain/wallet/models"
	"github.com/DevShuxat/eater-service/src/domain/wallet/repositories"
	"gorm.io/gorm"
)

const (
	tableWallet = "wallet.wallet"
)

type walletRepo struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) repositories.WalletRepository {
	return &walletRepo{
		db: db,
	}
}

func (r *walletRepo) SaveCard(ctx context.Context, card *models.PaymentCard) error {
	result := r.db.WithContext(ctx).Table(tableWallet).Create(card)
	if result.Error != nil {
		return result.Error
	}
	return nil
}


func (r *walletRepo) DeleteCard(ctx context.Context, cardID string) error {
 var wallet models.PaymentCard
 result := r.db.WithContext(ctx).Table(tableWallet).Delete(&wallet, "id = ?", cardID)
 if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *walletRepo) GetCard(ctx context.Context, cardID string) (*models.PaymentCard, error) {
	var wallet models.PaymentCard
	result := r.db.WithContext(ctx).Table(tableWallet).First(&wallet, "id = ?", cardID)
	if result.Error != nil {
		return nil, result.Error
	}
	return &wallet, nil
}

func (r *walletRepo) ListCardsByEater(ctx context.Context, eaterID string) ([]*models.PaymentCard, error) {
	var wallet []*models.PaymentCard
	result := r.db.WithContext(ctx).Table(tableWallet).Where("eater_id = ?", eaterID).Find(&wallet)
	if result.Error != nil {
		return nil, result.Error
	}
	return wallet, nil
}
