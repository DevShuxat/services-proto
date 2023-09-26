package services

import (
	"context"
	"errors"

	dtos "github.com/DevShuxat/eater-service/src/application/dtos/address"
	addressSvc "github.com/DevShuxat/eater-service/src/domain/address/services"
)

type AddressApplicationServices interface {
	SaveAddress(ctx context.Context, addressID, EaterID, Name string, Latitude, Longitude float64) (*dtos.CreateAddressResponse, error)
	UpdateAddress(ctx context.Context, addressID, EaterID, name, longitude, latitude string) (*dtos.UpdateAddressResponse, error)
	ListAddress(ctx context.Context, eaterID string) ([]*dtos.ListAddressResponse, error)
	GetAddress(ctx context.Context, ID string) (*dtos.GetAddressResponse, error)
	DeleteAddress(ctx context.Context,addressID string) (*dtos.DeleteAddressResponse, error)
}

type addressAppSvcImpl struct {
	addressSvc addressSvc.AddressService
}

func NewAddressApplicationServices(
	addressSvc addressSvc.AddressService,
) AddressApplicationServices {
	return &addressAppSvcImpl{
		addressSvc: addressSvc,
	}
}

func (s *addressAppSvcImpl) SaveAddress(ctx context.Context, addressID, EaterID, Name string, Latitude, Longitude float64) (*dtos.CreateAddressResponse, error) {
	if addressID == "" {
		return nil, errors.New("address is is required!")
	}

	if EaterID == "" {
		return nil, errors.New("EaterID is is required!")
	}

	if Name == "" {
		return nil, errors.New("name is is required!")
	}

	if Latitude == 0 {
		return nil, errors.New("Latitude is is required!")
	}

	if Longitude == 0 {
		return nil, errors.New("Longitude is is required!")
	}

	result, err := s.SaveAddress(ctx, addressID, EaterID, Name, Latitude, Longitude)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *addressAppSvcImpl) UpdateAddress(ctx context.Context, addressID, EaterID, Name string, Latitude, Longitude float64) (*dtos.CreateAddressResponse, error) {
	if addressID == "" {
		return nil, errors.New("address is is required!")
	}

	if EaterID == "" {
		return nil, errors.New("EaterID is is required!")
	}

	if Name == "" {
		return nil, errors.New("name is is required!")
	}

	if Latitude == 0 {
		return nil, errors.New("Latitude is is required!")
	}

	if Longitude == 0 {
		return nil, errors.New("Longitude is is required!")
	}

	result, err := s.UpdateAddress(ctx, addressID, EaterID, Name, Latitude, Longitude)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *addressAppSvcImpl) DeleteAddress(ctx context.Context,addressID string) (*dtos.DeleteAddressResponse, error) {
	if addressID == "" {
		return nil, errors.New("addressID is required!")
	}

	deleteAddressResponse := &dtos.DeleteAddressResponse{
		Success: false,
		Message: "",
	}

	return *deleteAddressResponse, nil
}

func (s *addressAppSvcImpl) GetAddress(ctx context.Context, addressID string) (*dtos.GetAddressResponse, error) {
	if addressID == "" {
		return nil, errors.New("addressID is is required!")
	}

	result, err := s.GetAddress(ctx, addressID)
	if err != nil {
		return nil, err
	}

	response := dtos.GetAddressResponse(*result)
	return &response, nil
}

func (s *addressAppSvcImpl) ListAddress(ctx context.Context, eaterID string) ([]*dtos.ListAddressResponse, error) {
	if eaterID == "" {
		return nil, errors.New("eaterID is is required!")
	}

	result, err := s.ListAddress(ctx, eaterID)
	if err != nil {
		return nil, err
	}

	return result, nil
}
