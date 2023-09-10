package dtos

type EaterSignupResponse struct {
	EaterID string `json:"eater"`
}

func NewEaterSignupResponse (eaterID string) *EaterSignupResponse {
	return &EaterSignupResponse{
		EaterID: eaterID,
	}
}
