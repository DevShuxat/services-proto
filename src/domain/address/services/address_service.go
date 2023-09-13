package services

import (
	"context"

	"github.com/DevShuxat/eater-service/src/domain/address/models"
	"github.com/DevShuxat/eater-service/src/domain/address/repositories"
)


type AddressService interface {
	SaveAddress(ctx context.Context, address *models.Address) (string error)
	UpdateAddress(ctx context.Context, address *models.Address) error
	GetAddress(ctx context.Context, addressID string) (*models.Address, error)
	DeleteAddress(ctx context.Context, addressID string) error
	ListAddressesByEaterID(ctx context.Context, eaterID string) ([]*models.Address, error)
}

type addressSvcImpl struct {
	addressRepo repositories.AddressRepository
}

func NewAddressService(addressRepo repositories.AddressRepository) AddressService {
    return &addressSvcImpl{
        addressRepo: addressRepo,
    }
}

func (s *addressSvcImpl) ListAddresses(ctx context.Context, eaterID string) ([]*models.Address, error) {
	addresses, err := s.addressRepo.ListAddressesByEaterID(ctx, eaterID)
	if err != nil {
		return nil, err
	}
	return addresses, nil
}