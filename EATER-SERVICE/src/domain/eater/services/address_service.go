package services

import (
	"github.com/DevShuxat/eater-service/EATER-SERVICE/src/domain/eater/service"
)

type AddressService struct {
	addresses map[string]*Address
}

func NewAddressService() *AddressService {
	return &AddressService{
		addresses: make(map[string]*Address),
	}
}

func (s *AddressService) GetAddress(id string) *Address {
	return s.addresses[id]
}

func (s *AddressService) AddAddress(addr *Address) {
	s.addresses[addr.ID] = addr
}

func (s *AddressService) UpdateAddress(id string, addr *Address) {
	s.addresses[id] = addr
}

func (s *AddressService) DeleteAddress(id string) {
	delete(s.addresses, id)
}