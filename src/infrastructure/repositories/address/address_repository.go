package repositories

import (
	"context"
	"errors"

	"github.com/DevShuxat/eater-service/src/domain/address/models"
	"github.com/DevShuxat/eater-service/src/domain/address/repositories"
	"gorm.io/gorm"
)

const (
	tableAddress = "address.address"
)

type NewAddressRepository struct {
	db *gorm.DB
}

// GetAddressById implements repositories.AddressRepository.
func (r *NewAddressRepository) GetAddressById(ctx context.Context, id string) (*models.Address, error) {
	var address models.Address
	err := r.db.WithContext(ctx).Table(tableAddress).Where("id = ?", id).First(&address).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("address not found")
		}
		return nil, err
	}
	return &address, nil
}

func (r *NewAddressRepository) DeleteAddress(ctx context.Context, id string) error {
	var address models.Address
	result := r.db.WithContext(ctx).Table(tableAddress).Delete(&address, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *NewAddressRepository) ListAddressByEaterId(ctx context.Context, eaterID string) (*models.Address, error) {
	var address models.Address
	result := r.db.WithContext(ctx).Table(tableAddress).Where(&address, "eater_id = ?", eaterID)
	if result.Error != nil {
		return nil, result.Error
	}
	return &address, nil
}

func NewGormAddressRepository(db *gorm.DB) *NewAddressRepository {
	return &NewAddressRepository{
		db: db,
	}
}

func (r *NewAddressRepository) CreateAddress(ctx context.Context, address *models.Address) error {
	err := r.db.WithContext(ctx).Create(&address).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *NewAddressRepository) UpdateAddress(ctx context.Context, address *models.Address) error {
	err := r.db.WithContext(ctx).Model(&models.Address{}).Where("id = ?", address.ID).Updates(&address).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *NewAddressRepository) GetAddressByID(ctx context.Context, addressID string) (*models.Address, error) {
	var address models.Address
	err := r.db.WithContext(ctx).Where("id = ?", addressID).First(&address).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("address not found")
		}
		return nil, err
	}
	return &address, nil
}

func (r *NewAddressRepository) WithTx(ctx context.Context, callback func(r repositories.AddressRepository) error) error {
	if err := r.db.Transaction(func(tx *gorm.DB) error {
		r := NewAddressRepository{
			db: tx,
		}
		if err := callback(&r); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return err
	}
	return nil
}
