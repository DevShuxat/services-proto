package services

import (
	"context"
	"errors"
	"time"

	"github.com/DevShuxat/eater-service/src/domain/address/models"
	"github.com/DevShuxat/eater-service/src/domain/address/repositories"
)


type AddressService interface {
	CreateAddress(ctx context.Context, addressID, EaterID, name string, Latitude, Longitude float64) (string error)
	UpdateAddress(ctx context.Context, Name, longitude, latitude string) (*models.Address, error)
	GetAddress(ctx context.Context, ID string) (*models.Address, error)
	DeleteAddress(ctx context.Context, ID string) (*models.Address, error)
	ListAddress(ctx context.Context, eaterID string) ([]*models.Address, error)
}

type addressSvcImpl struct {
	addressRepo repositories.AddressRepository
}

func (s *addressSvcImpl) GetAddress(ctx context.Context, ID string) (*models.Address, error) {
    address, err := s.addressRepo.GetAddress(ctx, ID)
    if err != nil {
        return nil, err
    }

    return &address, nil
}

func NewAddressService(addressRepo repositories.AddressRepository) AddressService {
    return &addressSvcImpl{
        addressRepo: addressRepo,
    }
}


func (s *addressSvcImpl) CreateAddress(ctx context.Context, addressID, EaterID, name string, Latitude, Longitude float64) (string, error) {
  location := &models.Location{
    Longitude: Longitude,
    Latitude:  Latitude,
  }

  address := &models.Address{
    ID:        addressID,
    EaterID:   EaterID,
    Name:      name,
    Location:  location,
    CreatedAt: time.Now().UTC(),
    UpdatedAt: time.Now().UTC(),
  }

  err := s.addressRepo.WithTx(ctx, func(r repositories.AddressRepository) error {
    if err := r.CreateAddress(ctx, address); err != nil {
      return err
    }
    return nil
  })
  if err != nil {
    return nil, err
  }
  return &address, nil
}

func (s *addressSvcImpl) UpdateAddress(ctx context.Context, address *models.Address) (*models.Address, error) {
    if address.ID == "" {
        return nil, errors.New("Address ID is required for updating")
    }

    existingAddress, err := s.addressRepo.GetAddress(ctx, address.ID)
    if err != nil {
        return nil, err
    }

    // Update the fields you want to change.
    existingAddress.Name = address.Name
    existingAddress.Location = address.Location
    existingAddress.UpdatedAt = time.Now().UTC()

    err = s.addressRepo.WithTx(ctx, func(r repositories.AddressRepository) error {
        if err := r.UpdateAddress(ctx, existingAddress); err != nil {
            return err
        }
        return nil
    })

    if err != nil {
        return nil, err
    }

    return &existingAddress, nil
}




