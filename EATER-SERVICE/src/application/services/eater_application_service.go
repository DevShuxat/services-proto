package services

import (
	"context"
)

type EaterApplicationService interface {
	SignupEater(ctx context.Context, phoneNumber string) (*dtos.EaterSignupResponse, error)
	ConfirmSMSCode(ctx context.Context, eaterID, smsCode string) (*dtos.ConfirmSMSCodeResponse, error)
	UpdateEaterProfile(ctx context.Context, eaterID, name, imageUrl string) (*dtos.UpdateEaterProfileResponse, error)
	GetEaterProfile(ctx context.Context, eaterID string) (*dtos.GetEaterProfileResponse, error)
	}