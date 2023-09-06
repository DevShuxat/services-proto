package services

import ("context")

type AddressApplicationServices interface {
	CreateAddress(ctx context.Context, name,longitude, latitude string) (*dtos.CreateAddressResponse, error)
	ListAddress(ctx context.Context, eaterID string) (*dtos.ListAddressResponse, error)
	UpdateAddress(ctx context.Context, name,longitude, latitude string) (*dtos.UpdateAddressResponse, error)
	DeleteAddress(ctx context.Context, id string) (*dtos.DeleteAddressResponse, error)
	GetAddress(ctx context.Context, id string) (*dtos.GetAddressResponse, error)
}