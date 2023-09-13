package services

import (
	"context"

	dtos "github.com/DevShuxat/eater-service/src/application/dtos/address"
)

type AddressApplicationServices interface {
	SaveAddress(ctx context.Context, name, longitude, latitude string) (*dtos.CreateAddressResponse, error)
	ListAddress(ctx context.Context, eaterID string) (*dtos.ListAddressResponse, error)
	UpdateAddress(ctx context.Context, name, longitude, latitude string) (*dtos.UpdateAddressResponse, error)
	DeleteAddress(ctx context.Context, ID string) (*dtos.DeleteAddressResponse, error)
	GetAddress(ctx context.Context, ID string) (*dtos.GetAddressResponse, error)
}
