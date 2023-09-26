package dtos

import "github.com/DevShuxat/eater-service/src/domain/address/models"

type ListAddressResponse struct {
	Addresses []*models.Address `json:"eater"`
}

func NewListAddressResponse(addresses []*models.Address) *ListAddressResponse {
	return &ListAddressResponse{
		Addresses: addresses,
	}
}
