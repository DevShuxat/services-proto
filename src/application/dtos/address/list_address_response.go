package dtos

import "github.com/DevShuxat/eater-service/src/domain/address/models"

type ListAddressResponse struct {
	Address *models.Address `json:"address"`
}

func NewListAddressResponse(address *models.Address) *GetAddressResponse {
	return &GetAddressResponse{
		Address: address,
	}
}
