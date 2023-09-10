package dtos

import "github.com/DevShuxat/eater-service/src/domain/address/models"

type ListAddressResponse struct {
	EaterID *models.Address `json:"address_by_eater_id"`
}

func NewListAddressResponse(EaterID string) *ListAddressResponse {
	return &ListAddressResponse{
		ListAddress: EaterID,
	}
}
