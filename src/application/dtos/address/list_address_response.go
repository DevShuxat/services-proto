package dtos

import "github.com/DevShuxat/eater-service/src/domain/address/models"

type ListAddressResponse struct {
	eaterID *models.EaterID `json:"address_by_eater_id"`
}

func NewListAddressResponse(eater_id string) *ListAddressResponse {
	return &ListAddressResponse{
		ListAddress: eater_id,
	}
}
