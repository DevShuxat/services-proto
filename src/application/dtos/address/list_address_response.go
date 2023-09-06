package dtos

import "github.com/DevShuxat/eater-service/src/domain/eater/models"

type ListAddressResponse struct {
	ListAddress *models.Address `json:"eater_id"`
}

func NewListAddressResponse(ListAddress *models.) *ListAddressResponse {
	return &ListAddressResponse{
		listAddress: *models.ListAddress,
	}
}
