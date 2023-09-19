package services

import (
	"context"
	"errors"
	"time"

	"github.com/DevShuxat/eater-service/src/domain/address/models"
	"github.com/DevShuxat/eater-service/src/domain/address/repositories"
)


type AddressService interface {
	SaveAddress(ctx context.Context, addressID, EaterID, Name string, Latitude, Longitude float64)  (*models.Address, error)
	UpdateAddress(ctx context.Context, address *models.Address) (error)
	DeleteAddress(ctx context.Context, addressID string) (error)
	GetAddress(ctx context.Context, addressID string) (*models.Address, error)
	ListAddressesByEater(ctx context.Context, eaterID string) ([]*models.Address, error)
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

func (s *addressSvcImpl) SaveAddress(ctx context.Context, addressID, EaterID, Name string, Latitude, Longitude float64) (*models.Address, error) {
  location := &models.Location{
    Longitude: Longitude,
    Latitude:  Latitude,
  }

  address := &models.Address{
    ID:        addressID,
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

func (s *addressSvcImpl) UpdateAddress(ctx context.Context, address *models.Address) (error) {
    if address.ID == "" {
        return errors.New("address ID is required for updating")

    }
    existingAddress, err := s.addressRepo.GetAddress(ctx, address.ID)
    if err != nil {
        return nil
    }

    existingAddress.Name = address.Name
    existingAddress.Location = address.Location
    existingAddress.UpdatedAt = time.Now().UTC()

        if err := s.addressRepo.UpdateAddress(ctx, existingAddress); err != nil {
            return err
        }
    return nil
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


func (s *addressSvcImpl) ListAddressesByEater(ctx context.Context, eaterID string) ([]*models.Address, error) {
    address, err := s.addressRepo.ListAddressesByEater(ctx, eaterID)
    if err != nil {
        return nil, err
    }
   
    return address, nil
}







 