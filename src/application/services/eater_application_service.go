package services

import (
	"context"
	"errors"
	"fmt"

	dtos "github.com/DevShuxat/eater-service/src/application/dtos/eater"
	EaterService "github.com/DevShuxat/eater-service/src/domain/eater/services"
	"github.com/DevShuxat/eater-service/src/infrastructure/jwt"
	"github.com/DevShuxat/eater-service/src/infrastructure/validator"
)

type EaterApplicationService interface {
	SignupEater(ctx context.Context, phoneNumber string) (*dtos.EaterSignupResponse, error)
	ConfirmSMSCode(ctx context.Context, eaterID, smsCode string) (*dtos.ConfirmSMSCodeResponse, error)
	UpdateEaterProfile(ctx context.Context, eaterID, name, imageUrl string) (*dtos.UpdateEaterProfileResponse, error)
	GetEaterProfile(ctx context.Context, eaterID string) (*dtos.GetEaterProfileResponse, error)
}

type eaterAppSvcImpl struct {
	eaterSvc EaterService.EaterService
	tokenSvc jwt.Service
}

func NewEaterApplicationService(
	eaterSvc EaterService.EaterService,
	tokenSvc jwt.Service,
) EaterApplicationService {
	return &eaterAppSvcImpl{
		eaterSvc: eaterSvc,
		tokenSvc: tokenSvc,
	}
}

func (s *eaterAppSvcImpl) SignupEater(ctx context.Context, phoneNumber string) (*dtos.EaterSignupResponse, error) {
	if !validator.ValidateUzPhoneNumber(phoneNumber) {
		return nil, errors.New("invalid phone number. ")
	}
	eaterId, err := s.eaterSvc.SignupEater(ctx, phoneNumber)
	if err != nil {
		return nil, err
	}

	return dtos.NewEaterSignupResponse(eaterId), nil
}

func (s *eaterAppSvcImpl) ConfirmSMSCode(ctx context.Context, eaterID, smsCode string) (*dtos.ConfirmSMSCodeResponse, error) {

	if eaterID == "" {
		return nil, fmt.Errorf("Invalid or missing eater_id: %s", eaterID)
	}

	if smsCode == "" {
		return nil, fmt.Errorf("Invalid or missing sms_code: %s", smsCode)
	}
	profile, err := s.eaterSvc.ConfirmSMSCode(ctx, eaterID, smsCode)
	if err != nil {
		return nil, err
	}

	token, err := s.tokenSvc.CreateToken(ctx, profile.EaterID)
	if err != nil {
		return nil, err
	}

	response := dtos.ConfirmSMSCodeResponse{
		Token:   token,
		Profile: profile,
	}
	return &response, nil
}

func (s *eaterAppSvcImpl) UpdateEaterProfile(ctx context.Context, eaterID, name, imageUrl string) (*dtos.UpdateEaterProfileResponse, error) {

	if eaterID == "" {
		return nil, fmt.Errorf("Invalid or missing eater_id: %s", eaterID)
	}

	if name == "" {
		return nil, fmt.Errorf("Invalid or missing name: %s", name)
	}

	profile, err := s.eaterSvc.UpdateEaterProfile(ctx, eaterID, name, imageUrl)
	if err != nil {
		return nil, err
	}

	return dtos.NewUpdateEaterProfileResponse(profile), nil
}

func (s *eaterAppSvcImpl) GetEaterProfile(ctx context.Context, eaterID string) (*dtos.GetEaterProfileResponse, error) {

	if eaterID == "" {
		return nil, fmt.Errorf("Invalid or missing eater_id: %s", eaterID)
	}

	profile, err := s.eaterSvc.GetEaterProfile(ctx, eaterID)
	if err != nil {
		return nil, err
	}

	return dtos.NewGetEaterProfileResponse(profile), nil
}
