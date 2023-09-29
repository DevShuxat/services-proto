package dtos

import (
	"github.com/DevShuxat/eater-service/src/domain/eater/models"
)
type ConfirmSMSCodeResponse struct {
	Token   string               `json:"token"`
	Profile *models.EaterProfile `json:"profile"`
}

func NewConfirmSMSCodeResponse(token string, profile *models.EaterProfile) *ConfirmSMSCodeResponse {
	return &ConfirmSMSCodeResponse{
		Token:   token,
		Profile: profile,
	}
}
