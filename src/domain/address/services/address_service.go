package services

import (
	"context"
	"errors"
	"time"

	"github.com/DevShuxat/eater-service/src/domain/address/models"
	"github.com/DevShuxat/eater-service/src/domain/address/repositories"
	"github.com/DevShuxat/eater-service/src/infrastructure/rand"
)


type AddressService interface {
	SaveAddress(ctx context.Context,  EaterID, Name string, Latitude, Longitude float64)  (*models.Address, error)
	UpdateAddress(ctx context.Context, addressId string, name string, long, lat float64) (*models.Address, error)
	DeleteAddress(ctx context.Context, addressID string) (error)
	GetAddress(ctx context.Context, addressID string) (*models.Address, error)
	ListAddressByEater(ctx context.Context, eaterID string, sort string, page, pageSize int) ([]*models.Address, error)
}

type addressSvcImpl struct {
	addressRepo repositories.AddressRepository
}

func NewAddressService(addressRepo repositories.AddressRepository) AddressService {
    return &addressSvcImpl{
        addressRepo: addressRepo,
    }
}

func (s *addressSvcImpl) GetAddress(ctx context.Context, addressID string) (*models.Address, error) {
    address, err := s.addressRepo.GetAddress(ctx, addressID)
    if err != nil {
        return nil, err
    }

    return address, nil
}

func (s *addressSvcImpl) SaveAddress(ctx context.Context, EaterID, Name string, Latitude, Longitude float64) (*models.Address, error) {
  location := &models.Location{
    Longitude: Longitude,
    Latitude:  Latitude,
  }

  address := &models.Address{
    ID:        rand.UUID(),
    EaterID:   EaterID,
    Name:      Name,
    Location:  location,
    CreatedAt: time.Now().UTC(),
    UpdatedAt: time.Now().UTC(),
  }

  err := s.addressRepo.SaveAddress(ctx, address)
    if err != nil {
    return nil, err
    }
    return address, nil
}

func (s *addressSvcImpl) UpdateAddress(ctx context.Context, addressId string, name string, long, lat float64) (*models.Address, error) {
   location := &models.Location{
		Longitude: long,
		Latitude:  lat,
	}

	address := &models.Address{
		ID:        addressId,
		Name:      name,
		Location:  location,
		UpdatedAt: time.Now().UTC(),
	}

	err := s.addressRepo.UpdateAddress(ctx, address)

	if err != nil {
		return nil, err
	}

	return address, nil
}

func (s *addressSvcImpl) DeleteAddress(ctx context.Context, addressID string) error {
    existingAddress, err := s.GetAddress(ctx, addressID)
    if err != nil {
        return err
    }
    if existingAddress == nil {
		return errors.New("eater profile not found")
	}
    return s.addressRepo.DeleteAddress(ctx, existingAddress.ID)
}


func (s *addressSvcImpl) ListAddressByEater(ctx context.Context, eaterID string, sort string, page, pageSize int) ([]*models.Address, error) {

	addresses, err := s.addressRepo.ListAddressByEater(ctx, eaterID, sort, page, pageSize)
	if err != nil {
		return nil, err
	}

	return addresses, nil
}







 