
package repositories

import (
	"context"
	"github.com/DevShuxat/eater-service/src/domain/address/models"
)

type AddressRepository interface {
	WithTx(ctx context.Context, f func(r AddressRepository) error) error
	CreateAddress(ctx context.Context, address *models.Address) error
	UpdateAddress(ctx context.Context, address *models.Address) error
	DeleteAddress(ctx context.Context, id string) error
	GetAddressById(ctx context.Context, id string) (*models.Address, error)
	ListAddressByEaterId(ctx context.Context, eaterID string) (*models.Address, error)
}
