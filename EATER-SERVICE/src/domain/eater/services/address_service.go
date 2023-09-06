package services

import (
	"context"

	"github.com/DevShuxat/eater-service/EATER-SERVICE/src/domain/eater/models"
	"github.com/DevShuxat/eater-service/EATER-SERVICE/src/domain/eater/repositories"
)


type AddressService interface {
	CreateAddress(ctx context.Context, name, longitude, latitude string) (string, error)
	UpdateAddress(ctx context.Context, name, longitude, latitude string) (*models.Address, error)
	GetAddress(ctx context.Context, id string) (*models.Address, error)
	DeleteAddress(ctx context.Context, id string) (*models.Address, error)
	ListAddress(ctx context.Context, eaterID  string) (*models.Address, error)
}


func NewAddressService(
	addressRepo *repositories.AddressRepository,
) AddressService {
	addresses: make(map[string]model * address),
}

func (s *AddressService) GetAddress(id string) *address {
	return s.addresses[id]
}

func (s *AddressService) UpdateAddress(id string, addr *Address) {
	s.addresses[id] = addr
}

func (s *AddressService) DeleteAddress(id string) {
	delete(s.addresses, id)
}
