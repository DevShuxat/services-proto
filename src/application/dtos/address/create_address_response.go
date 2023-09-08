package dtos

type CreateAddressResponse struct {
	Address string `json:"address"`
}

func NewCreateAddressResponse(address string) *CreateAddressResponse {
	return &CreateAddressResponse{Address: address}
}
