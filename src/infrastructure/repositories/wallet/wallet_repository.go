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

type walletRepoImpl struct {
	db *gorm.DB
}

func (r *walletRepoImpl) AddCard(ctx context.Context, CardToken string, Number string) (*models.PaymentCard, error) {
	newCard := &models.PaymentCard{
		CardToken: CardToken,
		Number:    Number,
	}

	result := r.db.WithContext(ctx).Table(tableWallet).Create(newCard)
	if result.Error != nil {
		return nil, result.Error
	}

	return newCard, nil
}


func (r *walletRepoImpl) DeleteCard(ctx context.Context, ID string) error {
 var wallet *models.PaymentCard
 result := r.db.WithContext(ctx).Table(tableWallet).Delete(&wallet, "id = ?", ID)
 if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *walletRepoImpl) GetCard(ctx context.Context, ID string) (*models.PaymentCard, error) {
	var wallet *models.PaymentCard
	result := r.db.WithContext(ctx).Table(tableWallet).First(&wallet, "id = ?", ID)
	if result.Error != nil {
		return result.Error
	}
	return &wallet, 
}

func (r *walletRepoImpl) ListCard(ctx context.Context, eaterID string) (*models.PaymentCard, error) {
	var wallet models.PaymentCard
	result := r.db.WithContext(ctx).Table(tableWallet).Where("eater_id = ?", eaterID).Find(&wallet)
	if result.Error {
		return nil, result.Error
	}
	return &wallet, nil
}

func NewWalletRepository(db *gorm.DB) *walletRepoImpl {
	return &walletRepoImpl{db: db}
}

func (r *walletRepoImpl) WithTx(ctx context.Context, callback func(r repositories.WalletRepository) error) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		r := walletRepoImpl{
			db: tx,
		}
		if err := callback(&r); err != nil {
			return err
		}
		return nil
	})
}