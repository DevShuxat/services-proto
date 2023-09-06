package dtos

import "github.com/DevShuxat/eater-service/src/domain/address/models"


type UpdateAddressResponse struct {
	Address *models.Address `json:"address"`
}

func NewUpdateAddress(address *models.Address) *UpdateAddressResponse {
	return &UpdateAddressResponse{
	Address: address,
	}
}
