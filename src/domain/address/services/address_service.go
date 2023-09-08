package services

import (
	"context"
	"github.com/google/uuid"
	
	"github.com/DevShuxat/eater-service/src/domain/address/models"
	"github.com/DevShuxat/eater-service/src/domain/eater/models"
	"github.com/DevShuxat/eater-service/src/domain/eater/repositories"
)


type AddressService interface {
	CreateAddress(ctx context.Context, name, longitude, latitude string) (string, error)
	UpdateAddress(ctx context.Context, name, longitude, latitude string) (*models.Address, error)
	GetAddress(ctx context.Context, id string) (*models.Address, error)
	DeleteAddress(ctx context.Context, id string) (*models.Address, error)
	ListAddress(ctx context.Context, eaterID  string) (*models.Address, error)
}

func addressSvcImpl struct {
	addressRepo repositories.AddressRepository
}


func NewAddressService(
	var (
		AddressID = rand.UUID()
		EaterID = address.eater_id
		Name = address.name
		Location = address.location
		now  = time.Now().UTC()
	)

	adddress = models.Address{
		ID: AddressID,
		EaterID: EaterID,
		Name: Name,
		Location: Location,
		CreatedAt: now,
		UpdatedAt: now,
	}

 err := s.addresRepo.WithTx(ctx, func(r repositories.AddressRepository) error {
    if err := r.CreateAddress(ctx, &adddress); err != nil {
        return err
    }
})
 if err !=nil {
	return "", nil
 }

 func (s *AddressService) GetAddress(id string) *address {
	return s.addresses[id]
}

func (s *addressSvcImpl) UpdateAddress(ctx context.Context, address *models.Address) error {
	if address.ID == "" {
		return errors.New("Address ID is required for updating")
	}

	existingAddress, err := s.addressRepo.GetAddressByID(ctx, address.ID)
	if err != nil {
		return err
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
		return err
	}

	return nil
}


func DeleteAddress(addressID string) (*dtos.DeleteAddressResponse, error) {
	success := true 
	var message string 
	deleteResponse := dtos.NewDeleteAddressResponse(success, message)
	return deleteResponse, nil
}
