package dtos


type DeleteAddressResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message,omitempty"`
}

func NewDeleteAddressResponse(success bool, message string) *DeleteAddressResponse {
	return &DeleteAddressResponse{
		Success: success,
		Message: message,
	}
}
