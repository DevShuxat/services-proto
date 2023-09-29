package repositories

import (
	"context"

	"github.com/DevShuxat/eater-service/src/domain/address/models"
)

type AddressRepository interface {
	SaveAddress(ctx context.Context, address *models.Address) error
	UpdateAddress(ctx context.Context, address *models.Address) error
	DeleteAddress(ctx context.Context, addressID string) error
	GetAddress(ctx context.Context, addressID string) (*models.Address, error)
	ListAddressByEater(ctx context.Context, eaterID string, sort string, page, pageSize int) ([]*models.Address, error)
}
