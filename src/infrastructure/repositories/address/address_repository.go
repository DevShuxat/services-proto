package repositories

import (
	"context"

	"github.com/DevShuxat/eater-service/src/domain/address/models"
	"github.com/DevShuxat/eater-service/src/domain/address/repositories"
	"gorm.io/gorm"
)

const (
	tableAddress = "eater.addresses"
)

type addressSvcImp struct {
	db *gorm.DB
}

func NewAddressRepository(db *gorm.DB) repositories.AddressRepository {
	return &addressSvcImp{
		db: db,
	}
}

func (r *addressSvcImp) SaveAddress(ctx context.Context, address *models.Address) error {
	err := r.db.WithContext(ctx).Create(&address).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *addressSvcImp) UpdateAddress(ctx context.Context, address *models.Address) error {
	err := r.db.WithContext(ctx).Table(tableAddress).Where("id = ?", address.ID).Updates(&address).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *addressSvcImp) DeleteAddress(ctx context.Context, addressID string) error {
	var address models.Address
	result := r.db.WithContext(ctx).Table(tableAddress).Delete(&address, "id = ?", addressID)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *addressSvcImp) GetAddress(ctx context.Context, addressID string) (*models.Address, error) {
	var address models.Address
	result := r.db.WithContext(ctx).
		Table(tableAddress).
		First(&address, addressID)
	if result.Error != nil {
		return nil, result.Error
	}
	return &address, nil
}

func (r *addressSvcImp) ListAddressByEater(ctx context.Context, eaterID string, sort string, page, pageSize int) ([]*models.Address, error) {
	var address []*models.Address
	result := r.db.WithContext(ctx).Table(tableAddress).Where(&address, "eater_id = ?", eaterID).Find(&address)

	if result.Error != nil {
		return nil, result.Error
	}
	return address, nil
}
