package repositories

import (
	"context"

	"github.com/DevShuxat/eater-service/src/domain/address/models"
)

type AddressRepository interface {
	SaveAddress(ctx context.Context, addressID, EaterID, Name string, Latitude, Longitude float64) (string, error)
	DeleteAddress(ctx context.Context, ID string) (*models.Address, error)
	UpdateAddress(ctx context.Context, Name, longitude, latitude string) (models.Address, error)
	GetAddress(ctx context.Context, ID string) (models.Address, error)
	ListAddressesByEaterID(ctx context.Context, eaterID string) ([]*models.Address, error)
}
